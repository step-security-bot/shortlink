package pgx

import (
	"context"
	"fmt"
	"reflect"
	"time"

	errors "golang.org/x/xerrors"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgproto3/v2"
	"github.com/jackc/pgtype"
)

// Rows is the result set returned from *Conn.Query. Rows must be closed before
// the *Conn can be used again. Rows are closed by explicitly calling Close(),
// calling Next() until it returns false, or when a fatal error occurs.
//
// Rows is an interface instead of a struct to allow tests to mock Query. However,
// adding a method to an interface is technically a breaking change. Because of this
// the Rows interface is partially excluded from semantic version requirements.
// Methods will not be removed or changed, but new methods may be added.
type Rows interface {
	// Close closes the rows, making the connection ready for use again. It is safe
	// to call Close after rows is already closed.
	Close()

	// Err returns any error that occurred while reading.
	Err() error

	// CommandTag returns the command tag from this query. It is only available after Rows is closed.
	CommandTag() pgconn.CommandTag

	FieldDescriptions() []pgproto3.FieldDescription

	// Next prepares the next row for reading. It returns true if there is another
	// row and false if no more rows are available. It automatically closes rows
	// when all rows are read.
	Next() bool

	// Scan reads the values from the current row into dest values positionally.
	// dest can include pointers to core types, values implementing the Scanner
	// interface, []byte, and nil. []byte will skip the decoding process and directly
	// copy the raw bytes received from PostgreSQL. nil will skip the value entirely.
	Scan(dest ...interface{}) error

	// Values returns the decoded row values.
	Values() ([]interface{}, error)

	// RawValues returns the unparsed bytes of the row values. The returned [][]byte is only valid until the next Next
	// call or the Rows is closed. However, the underlying byte data is safe to retain a reference to and mutate.
	RawValues() [][]byte
}

// Row is a convenience wrapper over Rows that is returned by QueryRow.
//
// Row is an interface instead of a struct to allow tests to mock QueryRow. However,
// adding a method to an interface is technically a breaking change. Because of this
// the Row interface is partially excluded from semantic version requirements.
// Methods will not be removed or changed, but new methods may be added.
type Row interface {
	// Scan works the same as Rows. with the following exceptions. If no
	// rows were found it returns ErrNoRows. If multiple rows are returned it
	// ignores all but the first.
	Scan(dest ...interface{}) error
}

// connRow implements the Row interface for Conn.QueryRow.
type connRow connRows

func (r *connRow) Scan(dest ...interface{}) (err error) {
	rows := (*connRows)(r)

	if rows.Err() != nil {
		return rows.Err()
	}

	if !rows.Next() {
		if rows.Err() == nil {
			return ErrNoRows
		}
		return rows.Err()
	}

	rows.Scan(dest...)
	rows.Close()
	return rows.Err()
}

type rowLog interface {
	shouldLog(lvl LogLevel) bool
	log(ctx context.Context, lvl LogLevel, msg string, data map[string]interface{})
}

// connRows implements the Rows interface for Conn.Query.
type connRows struct {
	ctx        context.Context
	logger     rowLog
	connInfo   *pgtype.ConnInfo
	values     [][]byte
	rowCount   int
	err        error
	commandTag pgconn.CommandTag
	startTime  time.Time
	sql        string
	args       []interface{}
	closed     bool

	resultReader      *pgconn.ResultReader
	multiResultReader *pgconn.MultiResultReader
}

func (rows *connRows) FieldDescriptions() []pgproto3.FieldDescription {
	return rows.resultReader.FieldDescriptions()
}

func (rows *connRows) Close() {
	if rows.closed {
		return
	}

	rows.closed = true

	if rows.resultReader != nil {
		var closeErr error
		rows.commandTag, closeErr = rows.resultReader.Close()
		if rows.err == nil {
			rows.err = closeErr
		}
	}

	if rows.multiResultReader != nil {
		closeErr := rows.multiResultReader.Close()
		if rows.err == nil {
			rows.err = closeErr
		}
	}

	if rows.logger != nil {
		if rows.err == nil {
			if rows.logger.shouldLog(LogLevelInfo) {
				endTime := time.Now()
				rows.logger.log(rows.ctx, LogLevelInfo, "Query", map[string]interface{}{"sql": rows.sql, "args": logQueryArgs(rows.args), "time": endTime.Sub(rows.startTime), "rowCount": rows.rowCount})
			}
		} else if rows.logger.shouldLog(LogLevelError) {
			rows.logger.log(rows.ctx, LogLevelError, "Query", map[string]interface{}{"sql": rows.sql, "args": logQueryArgs(rows.args)})
		}
	}
}

func (rows *connRows) CommandTag() pgconn.CommandTag {
	return rows.commandTag
}

func (rows *connRows) Err() error {
	return rows.err
}

// fatal signals an error occurred after the query was sent to the server. It
// closes the rows automatically.
func (rows *connRows) fatal(err error) {
	if rows.err != nil {
		return
	}

	rows.err = err
	rows.Close()
}

func (rows *connRows) Next() bool {
	if rows.closed {
		return false
	}

	if rows.resultReader.NextRow() {
		rows.rowCount++
		rows.values = rows.resultReader.Values()
		return true
	} else {
		rows.Close()
		return false
	}
}

func (rows *connRows) Scan(dest ...interface{}) error {
	err := ScanRow(rows.connInfo, rows.FieldDescriptions(), rows.values, dest...)
	if err != nil {
		rows.fatal(err)
		return err
	}

	return nil
}

func (rows *connRows) Values() ([]interface{}, error) {
	if rows.closed {
		return nil, errors.New("rows is closed")
	}

	values := make([]interface{}, 0, len(rows.FieldDescriptions()))

	for i := range rows.FieldDescriptions() {
		buf := rows.values[i]
		fd := &rows.FieldDescriptions()[i]

		if buf == nil {
			values = append(values, nil)
			continue
		}

		if dt, ok := rows.connInfo.DataTypeForOID(fd.DataTypeOID); ok {
			value := reflect.New(reflect.ValueOf(dt.Value).Elem().Type()).Interface().(pgtype.Value)

			switch fd.Format {
			case TextFormatCode:
				decoder, ok := value.(pgtype.TextDecoder)
				if !ok {
					decoder = &pgtype.GenericText{}
				}
				err := decoder.DecodeText(rows.connInfo, buf)
				if err != nil {
					rows.fatal(err)
				}
				values = append(values, decoder.(pgtype.Value).Get())
			case BinaryFormatCode:
				decoder, ok := value.(pgtype.BinaryDecoder)
				if !ok {
					decoder = &pgtype.GenericBinary{}
				}
				err := decoder.DecodeBinary(rows.connInfo, buf)
				if err != nil {
					rows.fatal(err)
				}
				values = append(values, value.Get())
			default:
				rows.fatal(errors.New("Unknown format code"))
			}
		} else {
			switch fd.Format {
			case TextFormatCode:
				decoder := &pgtype.GenericText{}
				err := decoder.DecodeText(rows.connInfo, buf)
				if err != nil {
					rows.fatal(err)
				}
				values = append(values, decoder.Get())
			case BinaryFormatCode:
				decoder := &pgtype.GenericBinary{}
				err := decoder.DecodeBinary(rows.connInfo, buf)
				if err != nil {
					rows.fatal(err)
				}
				values = append(values, decoder.Get())
			default:
				rows.fatal(errors.New("Unknown format code"))
			}
		}

		if rows.Err() != nil {
			return nil, rows.Err()
		}
	}

	return values, rows.Err()
}

func (rows *connRows) RawValues() [][]byte {
	return rows.values
}

type scanArgError struct {
	col int
	err error
}

func (e scanArgError) Error() string {
	return fmt.Sprintf("can't scan into dest[%d]: %v", e.col, e.err)
}

// ScanRow decodes raw row data into dest. This is a low level function used internally to to implement the Rows
// interface Scan method. It can be used to scan rows read from the lower level pgconn interface.
//
// connInfo - OID to Go type mapping.
// fieldDescriptions - OID and format of values
// values - the raw data as returned from the PostgreSQL server
// dest - the destination that values will be decoded into
func ScanRow(connInfo *pgtype.ConnInfo, fieldDescriptions []pgproto3.FieldDescription, values [][]byte, dest ...interface{}) error {
	if len(fieldDescriptions) != len(values) {
		return errors.Errorf("number of field descriptions must equal number of values, got %d and %d", len(fieldDescriptions), len(values))
	}
	if len(fieldDescriptions) != len(dest) {
		return errors.Errorf("number of field descriptions must equal number of destinations, got %d and %d", len(fieldDescriptions), len(dest))
	}

	for i, d := range dest {
		if d == nil {
			continue
		}

		err := connInfo.Scan(fieldDescriptions[i].DataTypeOID, fieldDescriptions[i].Format, values[i], d)
		if err != nil {
			return scanArgError{col: i, err: err}
		}
	}

	return nil
}
