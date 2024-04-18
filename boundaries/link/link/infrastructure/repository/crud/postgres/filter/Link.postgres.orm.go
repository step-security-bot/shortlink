// Code generated by protoc-gen-go-orm. DO NOT EDIT.
// versions:
// - protoc-gen-go-orm v1.6.0
// - protoc             (unknown)
// source: domain/link/v1/link.proto

package filter

import (
	"strings"
	"github.com/Masterminds/squirrel"
)

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
		if f.FieldMask.StartsWith != "" {
			query = query.Where("fieldmask LIKE '%' || ?", f.FieldMask.StartsWith)
		}
		if f.FieldMask.EndsWith != "" {
			query = query.Where("fieldmask LIKE ? || '%'", f.FieldMask.EndsWith)
		}
		if len(f.FieldMask.Contains) > 0 {
			containsQueries := []string{}
			containsArgs := []interface{}{}
			for _, v := range f.FieldMask.Contains {
				if v != "" {
					containsQueries = append(containsQueries, "fieldmask LIKE ?")
					containsArgs = append(containsArgs, "%"+v+"%")
				}
			}
			if len(containsQueries) > 0 {
				query = query.Where("("+strings.Join(containsQueries, " OR ")+")", containsArgs...)
			}
		}
		if len(f.FieldMask.NotContains) > 0 {
			notContainsQueries := []string{}
			notContainsArgs := []interface{}{}
			for _, v := range f.FieldMask.NotContains {
				if v != "" {
					notContainsQueries = append(notContainsQueries, "fieldmask NOT LIKE ?")
					notContainsArgs = append(notContainsArgs, "%"+v+"%")
				}
			}
			if len(notContainsQueries) > 0 {
				query = query.Where("("+strings.Join(notContainsQueries, " OR ")+")", notContainsArgs...)
			}
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
		if f.Url.StartsWith != "" {
			query = query.Where("url LIKE '%' || ?", f.Url.StartsWith)
		}
		if f.Url.EndsWith != "" {
			query = query.Where("url LIKE ? || '%'", f.Url.EndsWith)
		}
		if len(f.Url.Contains) > 0 {
			containsQueries := []string{}
			containsArgs := []interface{}{}
			for _, v := range f.Url.Contains {
				if v != "" {
					containsQueries = append(containsQueries, "url LIKE ?")
					containsArgs = append(containsArgs, "%"+v+"%")
				}
			}
			if len(containsQueries) > 0 {
				query = query.Where("("+strings.Join(containsQueries, " OR ")+")", containsArgs...)
			}
		}
		if len(f.Url.NotContains) > 0 {
			notContainsQueries := []string{}
			notContainsArgs := []interface{}{}
			for _, v := range f.Url.NotContains {
				if v != "" {
					notContainsQueries = append(notContainsQueries, "url NOT LIKE ?")
					notContainsArgs = append(notContainsArgs, "%"+v+"%")
				}
			}
			if len(notContainsQueries) > 0 {
				query = query.Where("("+strings.Join(notContainsQueries, " OR ")+")", notContainsArgs...)
			}
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
		if f.Hash.StartsWith != "" {
			query = query.Where("hash LIKE '%' || ?", f.Hash.StartsWith)
		}
		if f.Hash.EndsWith != "" {
			query = query.Where("hash LIKE ? || '%'", f.Hash.EndsWith)
		}
		if len(f.Hash.Contains) > 0 {
			containsQueries := []string{}
			containsArgs := []interface{}{}
			for _, v := range f.Hash.Contains {
				if v != "" {
					containsQueries = append(containsQueries, "hash LIKE ?")
					containsArgs = append(containsArgs, "%"+v+"%")
				}
			}
			if len(containsQueries) > 0 {
				query = query.Where("("+strings.Join(containsQueries, " OR ")+")", containsArgs...)
			}
		}
		if len(f.Hash.NotContains) > 0 {
			notContainsQueries := []string{}
			notContainsArgs := []interface{}{}
			for _, v := range f.Hash.NotContains {
				if v != "" {
					notContainsQueries = append(notContainsQueries, "hash NOT LIKE ?")
					notContainsArgs = append(notContainsArgs, "%"+v+"%")
				}
			}
			if len(notContainsQueries) > 0 {
				query = query.Where("("+strings.Join(notContainsQueries, " OR ")+")", notContainsArgs...)
			}
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
		if f.Describe.StartsWith != "" {
			query = query.Where("describe LIKE '%' || ?", f.Describe.StartsWith)
		}
		if f.Describe.EndsWith != "" {
			query = query.Where("describe LIKE ? || '%'", f.Describe.EndsWith)
		}
		if len(f.Describe.Contains) > 0 {
			containsQueries := []string{}
			containsArgs := []interface{}{}
			for _, v := range f.Describe.Contains {
				if v != "" {
					containsQueries = append(containsQueries, "describe LIKE ?")
					containsArgs = append(containsArgs, "%"+v+"%")
				}
			}
			if len(containsQueries) > 0 {
				query = query.Where("("+strings.Join(containsQueries, " OR ")+")", containsArgs...)
			}
		}
		if len(f.Describe.NotContains) > 0 {
			notContainsQueries := []string{}
			notContainsArgs := []interface{}{}
			for _, v := range f.Describe.NotContains {
				if v != "" {
					notContainsQueries = append(notContainsQueries, "describe NOT LIKE ?")
					notContainsArgs = append(notContainsArgs, "%"+v+"%")
				}
			}
			if len(notContainsQueries) > 0 {
				query = query.Where("("+strings.Join(notContainsQueries, " OR ")+")", notContainsArgs...)
			}
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
		if f.CreatedAt.StartsWith != "" {
			query = query.Where("createdat LIKE '%' || ?", f.CreatedAt.StartsWith)
		}
		if f.CreatedAt.EndsWith != "" {
			query = query.Where("createdat LIKE ? || '%'", f.CreatedAt.EndsWith)
		}
		if len(f.CreatedAt.Contains) > 0 {
			containsQueries := []string{}
			containsArgs := []interface{}{}
			for _, v := range f.CreatedAt.Contains {
				if v != "" {
					containsQueries = append(containsQueries, "createdat LIKE ?")
					containsArgs = append(containsArgs, "%"+v+"%")
				}
			}
			if len(containsQueries) > 0 {
				query = query.Where("("+strings.Join(containsQueries, " OR ")+")", containsArgs...)
			}
		}
		if len(f.CreatedAt.NotContains) > 0 {
			notContainsQueries := []string{}
			notContainsArgs := []interface{}{}
			for _, v := range f.CreatedAt.NotContains {
				if v != "" {
					notContainsQueries = append(notContainsQueries, "createdat NOT LIKE ?")
					notContainsArgs = append(notContainsArgs, "%"+v+"%")
				}
			}
			if len(notContainsQueries) > 0 {
				query = query.Where("("+strings.Join(notContainsQueries, " OR ")+")", notContainsArgs...)
			}
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
		if f.UpdatedAt.StartsWith != "" {
			query = query.Where("updatedat LIKE '%' || ?", f.UpdatedAt.StartsWith)
		}
		if f.UpdatedAt.EndsWith != "" {
			query = query.Where("updatedat LIKE ? || '%'", f.UpdatedAt.EndsWith)
		}
		if len(f.UpdatedAt.Contains) > 0 {
			containsQueries := []string{}
			containsArgs := []interface{}{}
			for _, v := range f.UpdatedAt.Contains {
				if v != "" {
					containsQueries = append(containsQueries, "updatedat LIKE ?")
					containsArgs = append(containsArgs, "%"+v+"%")
				}
			}
			if len(containsQueries) > 0 {
				query = query.Where("("+strings.Join(containsQueries, " OR ")+")", containsArgs...)
			}
		}
		if len(f.UpdatedAt.NotContains) > 0 {
			notContainsQueries := []string{}
			notContainsArgs := []interface{}{}
			for _, v := range f.UpdatedAt.NotContains {
				if v != "" {
					notContainsQueries = append(notContainsQueries, "updatedat NOT LIKE ?")
					notContainsArgs = append(notContainsArgs, "%"+v+"%")
				}
			}
			if len(notContainsQueries) > 0 {
				query = query.Where("("+strings.Join(notContainsQueries, " OR ")+")", notContainsArgs...)
			}
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