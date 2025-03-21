// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package crud

import (
	"database/sql"
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

// Link list
type Link struct {
	ID        uuid.UUID
	Url       string
	Hash      string
	Describe  sql.NullString
	Json      json.RawMessage
	CreatedAt time.Time
	UpdatedAt time.Time
}
