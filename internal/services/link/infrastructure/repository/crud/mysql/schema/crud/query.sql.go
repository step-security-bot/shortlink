// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0
// source: query.sql

package crud

import (
	"context"
	"database/sql"
	"encoding/json"

	"github.com/google/uuid"
)

const createLink = `-- name: CreateLink :execresult
INSERT INTO links (id, url, hash, ` + "`" + `describe` + "`" + `, json)
  VALUES (?, ?, ?, ?, ?)
`

type CreateLinkParams struct {
	ID       uuid.UUID
	Url      string
	Hash     string
	Describe sql.NullString
	Json     json.RawMessage
}

func (q *Queries) CreateLink(ctx context.Context, arg CreateLinkParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, createLink,
		arg.ID,
		arg.Url,
		arg.Hash,
		arg.Describe,
		arg.Json,
	)
}

const deleteLink = `-- name: DeleteLink :exec
DELETE FROM links
  WHERE hash = ?
`

func (q *Queries) DeleteLink(ctx context.Context, hash string) error {
	_, err := q.db.ExecContext(ctx, deleteLink, hash)
	return err
}

const getLinkByHash = `-- name: GetLinkByHash :one
SELECT id, url, hash, ` + "`" + `describe` + "`" + `, json, created_at, updated_at FROM links
  WHERE hash = ?
`

func (q *Queries) GetLinkByHash(ctx context.Context, hash string) (Link, error) {
	row := q.db.QueryRowContext(ctx, getLinkByHash, hash)
	var i Link
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
SELECT id, url, hash, ` + "`" + `describe` + "`" + `, json, created_at, updated_at FROM links
`

func (q *Queries) GetLinks(ctx context.Context) ([]Link, error) {
	rows, err := q.db.QueryContext(ctx, getLinks)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Link
	for rows.Next() {
		var i Link
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
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateLink = `-- name: UpdateLink :execresult
UPDATE links
    SET url = ?, hash = ?, ` + "`" + `describe` + "`" + ` = ?, json = ?
WHERE id = ?
`

type UpdateLinkParams struct {
	Url      string
	Hash     string
	Describe sql.NullString
	Json     json.RawMessage
	ID       uuid.UUID
}

func (q *Queries) UpdateLink(ctx context.Context, arg UpdateLinkParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, updateLink,
		arg.Url,
		arg.Hash,
		arg.Describe,
		arg.Json,
		arg.ID,
	)
}
