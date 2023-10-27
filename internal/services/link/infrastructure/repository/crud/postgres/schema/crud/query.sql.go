// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0
// source: query.sql

package crud

import (
	"context"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgtype"
	v1 "github.com/shortlink-org/shortlink/internal/services/link/domain/link/v1"
)

const createLink = `-- name: CreateLink :execresult
INSERT INTO link.links (id, url, hash, describe, json)
VALUES ($1, $2, $3, $4, $5)
`

type CreateLinkParams struct {
	ID       pgtype.UUID
	Url      string
	Hash     string
	Describe pgtype.Text
	Json     v1.Link
}

func (q *Queries) CreateLink(ctx context.Context, arg CreateLinkParams) (pgconn.CommandTag, error) {
	return q.db.Exec(ctx, createLink,
		arg.ID,
		arg.Url,
		arg.Hash,
		arg.Describe,
		arg.Json,
	)
}

const deleteLink = `-- name: DeleteLink :exec
DELETE FROM link.links
WHERE hash = $1
`

func (q *Queries) DeleteLink(ctx context.Context, hash string) error {
	_, err := q.db.Exec(ctx, deleteLink, hash)
	return err
}

const getLinkByHash = `-- name: GetLinkByHash :one
SELECT id, url, hash, describe, json, created_at, updated_at FROM link.links
WHERE hash = $1
`

func (q *Queries) GetLinkByHash(ctx context.Context, hash string) (LinkLink, error) {
	row := q.db.QueryRow(ctx, getLinkByHash, hash)
	var i LinkLink
	err := row.Scan(
		&i.ID,
		&i.Url,
		&i.Hash,
		&i.Describe,
		&i.Json,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getLinks = `-- name: GetLinks :many
SELECT id, url, hash, describe, json, created_at, updated_at FROM link.links
`

func (q *Queries) GetLinks(ctx context.Context) ([]LinkLink, error) {
	rows, err := q.db.Query(ctx, getLinks)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []LinkLink
	for rows.Next() {
		var i LinkLink
		if err := rows.Scan(
			&i.ID,
			&i.Url,
			&i.Hash,
			&i.Describe,
			&i.Json,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateLink = `-- name: UpdateLink :execresult
UPDATE link.links
SET url = $1, hash = $2, describe = $3, json = $4
WHERE id = $5
`

type UpdateLinkParams struct {
	Url      string
	Hash     string
	Describe pgtype.Text
	Json     v1.Link
	ID       pgtype.UUID
}

func (q *Queries) UpdateLink(ctx context.Context, arg UpdateLinkParams) (pgconn.CommandTag, error) {
	return q.db.Exec(ctx, updateLink,
		arg.Url,
		arg.Hash,
		arg.Describe,
		arg.Json,
		arg.ID,
	)
}
