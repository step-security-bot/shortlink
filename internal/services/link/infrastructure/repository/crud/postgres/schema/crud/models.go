// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0

package crud

import (
	"github.com/jackc/pgx/v5/pgtype"
	v1 "github.com/shortlink-org/shortlink/internal/services/link/domain/link/v1"
)

type LinkLink struct {
	ID        pgtype.UUID
	Url       string
	Hash      string
	Describe  pgtype.Text
	Json      v1.Link
	CreatedAt pgtype.Timestamp
	UpdatedAt pgtype.Timestamp
}

// CQRS for links
type LinkLinkView struct {
	ID              pgtype.UUID
	Url             string
	Hash            string
	Describe        pgtype.Text
	CreatedAt       pgtype.Timestamp
	UpdatedAt       pgtype.Timestamp
	ImageUrl        pgtype.Text
	MetaDescription pgtype.Text
	MetaKeywords    pgtype.Text
}
