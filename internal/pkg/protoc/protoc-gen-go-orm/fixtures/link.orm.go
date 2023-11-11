// Code generated by protoc-gen-go-orm. DO NOT EDIT.
// versions:
// - protoc-gen-go-orm v1.0.0
// - protoc             v4.24.2
// source: fixtures/link.proto

package fixtures

import (
	"github.com/Masterminds/squirrel"
)

type FilterLink struct {
	FieldMask  *StringFilterInput `json:"fieldmask"`
	Url        *StringFilterInput `json:"url"`
	Hash       *StringFilterInput `json:"hash"`
	Describe   *StringFilterInput `json:"describe"`
	CreatedAt  *StringFilterInput `json:"createdat"`
	UpdatedAt  *StringFilterInput `json:"updatedat"`
	Pagination *Pagination        `json:"pagination,omitempty"`
}

func (f *FilterLink) BuildFilter(query squirrel.SelectBuilder) squirrel.SelectBuilder {
	if f.FieldMask != nil {
		if f.FieldMask.Eq != "" {
			query = query.Where("fieldmask = ?", f.FieldMask.Eq)
		}
		if f.FieldMask.Ne != "" {
			query = query.Where("fieldmask <> ?", f.FieldMask.Ne)
		}
		if f.FieldMask.Lt != "" {
			query = query.Where("fieldmask < ?", f.FieldMask.Lt)
		}
		if f.FieldMask.Le != "" {
			query = query.Where("fieldmask <= ?", f.FieldMask.Le)
		}
		if f.FieldMask.Gt != "" {
			query = query.Where("fieldmask > ?", f.FieldMask.Gt)
		}
		if f.FieldMask.Ge != "" {
			query = query.Where("fieldmask >= ?", f.FieldMask.Ge)
		}
		if f.FieldMask.Contains != "" {
			query = query.Where("fieldmask LIKE '%' || ?", f.FieldMask.Contains)
		}
		if f.FieldMask.NotContains != "" {
			query = query.Where("fieldmask NOT LIKE ?", f.FieldMask.NotContains)
		}
		if f.FieldMask.StartsWith != "" {
			query = query.Where("fieldmask LIKE '%' || ?", f.FieldMask.StartsWith)
		}
		if f.FieldMask.EndsWith != "" {
			query = query.Where("fieldmask LIKE ? || '%'", f.FieldMask.EndsWith)
		}
		if f.FieldMask.IsEmpty {
			query = query.Where("fieldmask = '' OR fieldmask IS NULL")
		}
		if f.FieldMask.IsNotEmpty {
			query = query.Where("fieldmask <> '' AND fieldmask IS NOT NULL")
		}
	}
	if f.Url != nil {
		if f.Url.Eq != "" {
			query = query.Where("url = ?", f.Url.Eq)
		}
		if f.Url.Ne != "" {
			query = query.Where("url <> ?", f.Url.Ne)
		}
		if f.Url.Lt != "" {
			query = query.Where("url < ?", f.Url.Lt)
		}
		if f.Url.Le != "" {
			query = query.Where("url <= ?", f.Url.Le)
		}
		if f.Url.Gt != "" {
			query = query.Where("url > ?", f.Url.Gt)
		}
		if f.Url.Ge != "" {
			query = query.Where("url >= ?", f.Url.Ge)
		}
		if f.Url.Contains != "" {
			query = query.Where("url LIKE '%' || ?", f.Url.Contains)
		}
		if f.Url.NotContains != "" {
			query = query.Where("url NOT LIKE ?", f.Url.NotContains)
		}
		if f.Url.StartsWith != "" {
			query = query.Where("url LIKE '%' || ?", f.Url.StartsWith)
		}
		if f.Url.EndsWith != "" {
			query = query.Where("url LIKE ? || '%'", f.Url.EndsWith)
		}
		if f.Url.IsEmpty {
			query = query.Where("url = '' OR url IS NULL")
		}
		if f.Url.IsNotEmpty {
			query = query.Where("url <> '' AND url IS NOT NULL")
		}
	}
	if f.Hash != nil {
		if f.Hash.Eq != "" {
			query = query.Where("hash = ?", f.Hash.Eq)
		}
		if f.Hash.Ne != "" {
			query = query.Where("hash <> ?", f.Hash.Ne)
		}
		if f.Hash.Lt != "" {
			query = query.Where("hash < ?", f.Hash.Lt)
		}
		if f.Hash.Le != "" {
			query = query.Where("hash <= ?", f.Hash.Le)
		}
		if f.Hash.Gt != "" {
			query = query.Where("hash > ?", f.Hash.Gt)
		}
		if f.Hash.Ge != "" {
			query = query.Where("hash >= ?", f.Hash.Ge)
		}
		if f.Hash.Contains != "" {
			query = query.Where("hash LIKE '%' || ?", f.Hash.Contains)
		}
		if f.Hash.NotContains != "" {
			query = query.Where("hash NOT LIKE ?", f.Hash.NotContains)
		}
		if f.Hash.StartsWith != "" {
			query = query.Where("hash LIKE '%' || ?", f.Hash.StartsWith)
		}
		if f.Hash.EndsWith != "" {
			query = query.Where("hash LIKE ? || '%'", f.Hash.EndsWith)
		}
		if f.Hash.IsEmpty {
			query = query.Where("hash = '' OR hash IS NULL")
		}
		if f.Hash.IsNotEmpty {
			query = query.Where("hash <> '' AND hash IS NOT NULL")
		}
	}
	if f.Describe != nil {
		if f.Describe.Eq != "" {
			query = query.Where("describe = ?", f.Describe.Eq)
		}
		if f.Describe.Ne != "" {
			query = query.Where("describe <> ?", f.Describe.Ne)
		}
		if f.Describe.Lt != "" {
			query = query.Where("describe < ?", f.Describe.Lt)
		}
		if f.Describe.Le != "" {
			query = query.Where("describe <= ?", f.Describe.Le)
		}
		if f.Describe.Gt != "" {
			query = query.Where("describe > ?", f.Describe.Gt)
		}
		if f.Describe.Ge != "" {
			query = query.Where("describe >= ?", f.Describe.Ge)
		}
		if f.Describe.Contains != "" {
			query = query.Where("describe LIKE '%' || ?", f.Describe.Contains)
		}
		if f.Describe.NotContains != "" {
			query = query.Where("describe NOT LIKE ?", f.Describe.NotContains)
		}
		if f.Describe.StartsWith != "" {
			query = query.Where("describe LIKE '%' || ?", f.Describe.StartsWith)
		}
		if f.Describe.EndsWith != "" {
			query = query.Where("describe LIKE ? || '%'", f.Describe.EndsWith)
		}
		if f.Describe.IsEmpty {
			query = query.Where("describe = '' OR describe IS NULL")
		}
		if f.Describe.IsNotEmpty {
			query = query.Where("describe <> '' AND describe IS NOT NULL")
		}
	}
	if f.CreatedAt != nil {
		if f.CreatedAt.Eq != "" {
			query = query.Where("createdat = ?", f.CreatedAt.Eq)
		}
		if f.CreatedAt.Ne != "" {
			query = query.Where("createdat <> ?", f.CreatedAt.Ne)
		}
		if f.CreatedAt.Lt != "" {
			query = query.Where("createdat < ?", f.CreatedAt.Lt)
		}
		if f.CreatedAt.Le != "" {
			query = query.Where("createdat <= ?", f.CreatedAt.Le)
		}
		if f.CreatedAt.Gt != "" {
			query = query.Where("createdat > ?", f.CreatedAt.Gt)
		}
		if f.CreatedAt.Ge != "" {
			query = query.Where("createdat >= ?", f.CreatedAt.Ge)
		}
		if f.CreatedAt.Contains != "" {
			query = query.Where("createdat LIKE '%' || ?", f.CreatedAt.Contains)
		}
		if f.CreatedAt.NotContains != "" {
			query = query.Where("createdat NOT LIKE ?", f.CreatedAt.NotContains)
		}
		if f.CreatedAt.StartsWith != "" {
			query = query.Where("createdat LIKE '%' || ?", f.CreatedAt.StartsWith)
		}
		if f.CreatedAt.EndsWith != "" {
			query = query.Where("createdat LIKE ? || '%'", f.CreatedAt.EndsWith)
		}
		if f.CreatedAt.IsEmpty {
			query = query.Where("createdat = '' OR createdat IS NULL")
		}
		if f.CreatedAt.IsNotEmpty {
			query = query.Where("createdat <> '' AND createdat IS NOT NULL")
		}
	}
	if f.UpdatedAt != nil {
		if f.UpdatedAt.Eq != "" {
			query = query.Where("updatedat = ?", f.UpdatedAt.Eq)
		}
		if f.UpdatedAt.Ne != "" {
			query = query.Where("updatedat <> ?", f.UpdatedAt.Ne)
		}
		if f.UpdatedAt.Lt != "" {
			query = query.Where("updatedat < ?", f.UpdatedAt.Lt)
		}
		if f.UpdatedAt.Le != "" {
			query = query.Where("updatedat <= ?", f.UpdatedAt.Le)
		}
		if f.UpdatedAt.Gt != "" {
			query = query.Where("updatedat > ?", f.UpdatedAt.Gt)
		}
		if f.UpdatedAt.Ge != "" {
			query = query.Where("updatedat >= ?", f.UpdatedAt.Ge)
		}
		if f.UpdatedAt.Contains != "" {
			query = query.Where("updatedat LIKE '%' || ?", f.UpdatedAt.Contains)
		}
		if f.UpdatedAt.NotContains != "" {
			query = query.Where("updatedat NOT LIKE ?", f.UpdatedAt.NotContains)
		}
		if f.UpdatedAt.StartsWith != "" {
			query = query.Where("updatedat LIKE '%' || ?", f.UpdatedAt.StartsWith)
		}
		if f.UpdatedAt.EndsWith != "" {
			query = query.Where("updatedat LIKE ? || '%'", f.UpdatedAt.EndsWith)
		}
		if f.UpdatedAt.IsEmpty {
			query = query.Where("updatedat = '' OR updatedat IS NULL")
		}
		if f.UpdatedAt.IsNotEmpty {
			query = query.Where("updatedat <> '' AND updatedat IS NOT NULL")
		}
	}
	return query
}

type FilterLinks struct {
	Pagination *Pagination `json:"pagination,omitempty"`
}

func (f *FilterLinks) BuildFilter(query squirrel.SelectBuilder) squirrel.SelectBuilder {
	return query
}
