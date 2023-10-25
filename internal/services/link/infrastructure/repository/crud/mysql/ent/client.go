// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"
	"reflect"

	"github.com/google/uuid"
	"github.com/shortlink-org/shortlink/internal/services/link/infrastructure/repository/crud/mysql/ent/migrate"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"github.com/shortlink-org/shortlink/internal/services/link/infrastructure/repository/crud/mysql/ent/link"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Link is the client for interacting with the Link builders.
	Link *LinkClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}, inters: &inters{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Link = NewLinkClient(c.config)
}

type (
	// config is the configuration for the client and its builder.
	config struct {
		// driver used for executing database requests.
		driver dialect.Driver
		// debug enable a debug logging.
		debug bool
		// log used for logging on debug mode.
		log func(...any)
		// hooks to execute on mutations.
		hooks *hooks
		// interceptors to execute on queries.
		inters *inters
	}
	// Option function to configure the client.
	Option func(*config)
)

// options applies the options on the config object.
func (c *config) options(opts ...Option) {
	for _, opt := range opts {
		opt(c)
	}
	if c.debug {
		c.driver = dialect.Debug(c.driver, c.log)
	}
}

// Debug enables debug logging on the ent.Driver.
func Debug() Option {
	return func(c *config) {
		c.debug = true
	}
}

// Log sets the logging function for debug mode.
func Log(fn func(...any)) Option {
	return func(c *config) {
		c.log = fn
	}
}

// Driver configures the client driver.
func Driver(driver dialect.Driver) Option {
	return func(c *config) {
		c.driver = driver
	}
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// ErrTxStarted is returned when trying to start a new transaction from a transactional client.
var ErrTxStarted = errors.New("ent: cannot start a transaction within a transaction")

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, ErrTxStarted
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:    ctx,
		config: cfg,
		Link:   NewLinkClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		ctx:    ctx,
		config: cfg,
		Link:   NewLinkClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Link.
//		Query().
//		Count(ctx)
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.Link.Use(hooks...)
}

// Intercept adds the query interceptors to all the entity clients.
// In order to add interceptors to a specific client, call: `client.Node.Intercept(...)`.
func (c *Client) Intercept(interceptors ...Interceptor) {
	c.Link.Intercept(interceptors...)
}

// Mutate implements the ent.Mutator interface.
func (c *Client) Mutate(ctx context.Context, m Mutation) (Value, error) {
	switch m := m.(type) {
	case *LinkMutation:
		return c.Link.mutate(ctx, m)
	default:
		return nil, fmt.Errorf("ent: unknown mutation type %T", m)
	}
}

// LinkClient is a client for the Link schema.
type LinkClient struct {
	config
}

// NewLinkClient returns a client for the Link from the given config.
func NewLinkClient(c config) *LinkClient {
	return &LinkClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `link.Hooks(f(g(h())))`.
func (c *LinkClient) Use(hooks ...Hook) {
	c.hooks.Link = append(c.hooks.Link, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `link.Intercept(f(g(h())))`.
func (c *LinkClient) Intercept(interceptors ...Interceptor) {
	c.inters.Link = append(c.inters.Link, interceptors...)
}

// Create returns a builder for creating a Link entity.
func (c *LinkClient) Create() *LinkCreate {
	mutation := newLinkMutation(c.config, OpCreate)
	return &LinkCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Link entities.
func (c *LinkClient) CreateBulk(builders ...*LinkCreate) *LinkCreateBulk {
	return &LinkCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *LinkClient) MapCreateBulk(slice any, setFunc func(*LinkCreate, int)) *LinkCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &LinkCreateBulk{err: fmt.Errorf("calling to LinkClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*LinkCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &LinkCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Link.
func (c *LinkClient) Update() *LinkUpdate {
	mutation := newLinkMutation(c.config, OpUpdate)
	return &LinkUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *LinkClient) UpdateOne(l *Link) *LinkUpdateOne {
	mutation := newLinkMutation(c.config, OpUpdateOne, withLink(l))
	return &LinkUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *LinkClient) UpdateOneID(id uuid.UUID) *LinkUpdateOne {
	mutation := newLinkMutation(c.config, OpUpdateOne, withLinkID(id))
	return &LinkUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Link.
func (c *LinkClient) Delete() *LinkDelete {
	mutation := newLinkMutation(c.config, OpDelete)
	return &LinkDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *LinkClient) DeleteOne(l *Link) *LinkDeleteOne {
	return c.DeleteOneID(l.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *LinkClient) DeleteOneID(id uuid.UUID) *LinkDeleteOne {
	builder := c.Delete().Where(link.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &LinkDeleteOne{builder}
}

// Query returns a query builder for Link.
func (c *LinkClient) Query() *LinkQuery {
	return &LinkQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeLink},
		inters: c.Interceptors(),
	}
}

// Get returns a Link entity by its id.
func (c *LinkClient) Get(ctx context.Context, id uuid.UUID) (*Link, error) {
	return c.Query().Where(link.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *LinkClient) GetX(ctx context.Context, id uuid.UUID) *Link {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *LinkClient) Hooks() []Hook {
	return c.hooks.Link
}

// Interceptors returns the client interceptors.
func (c *LinkClient) Interceptors() []Interceptor {
	return c.inters.Link
}

func (c *LinkClient) mutate(ctx context.Context, m *LinkMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&LinkCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&LinkUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&LinkUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&LinkDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Link mutation op: %q", m.Op())
	}
}

// hooks and interceptors per client, for fast access.
type (
	hooks struct {
		Link []ent.Hook
	}
	inters struct {
		Link []ent.Interceptor
	}
)
