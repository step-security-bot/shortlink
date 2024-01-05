// Code generated by protoc-gen-go-orm. DO NOT EDIT.
// versions:
// - protoc-gen-go-orm v1.1.0
// - protoc             (unknown)
// source: infrastructure/rpc/cqrs/link/v1/link_command.proto

package v1

import (
	"strings"
	"github.com/Masterminds/squirrel"
	"go.mongodb.org/mongo-driver/bson"
)

type FilterAddRequest struct {
	Link       *StringFilterInput `json:"link"`
	Pagination *Pagination        `json:"pagination,omitempty"`
}

func (f *FilterAddRequest) BuildFilter(query squirrel.SelectBuilder) squirrel.SelectBuilder {
	if f.Link != nil {
		if f.Link.Eq != "" {
			query = query.Where("link = ?", f.Link.Eq)
		}
		if f.Link.Ne != "" {
			query = query.Where("link <> ?", f.Link.Ne)
		}
		if f.Link.Lt != "" {
			query = query.Where("link < ?", f.Link.Lt)
		}
		if f.Link.Le != "" {
			query = query.Where("link <= ?", f.Link.Le)
		}
		if f.Link.Gt != "" {
			query = query.Where("link > ?", f.Link.Gt)
		}
		if f.Link.Ge != "" {
			query = query.Where("link >= ?", f.Link.Ge)
		}
		if f.Link.StartsWith != "" {
			query = query.Where("link LIKE '%' || ?", f.Link.StartsWith)
		}
		if f.Link.EndsWith != "" {
			query = query.Where("link LIKE ? || '%'", f.Link.EndsWith)
		}
		if len(f.Link.Contains) > 0 {
			containsQueries := []string{}
			containsArgs := []interface{}{}
			for _, v := range f.Link.Contains {
				if v != "" {
					containsQueries = append(containsQueries, "link LIKE ?")
					containsArgs = append(containsArgs, "%"+v+"%")
				}
			}
			if len(containsQueries) > 0 {
				query = query.Where("("+strings.Join(containsQueries, " OR ")+")", containsArgs...)
			}
		}
		if len(f.Link.NotContains) > 0 {
			notContainsQueries := []string{}
			notContainsArgs := []interface{}{}
			for _, v := range f.Link.NotContains {
				if v != "" {
					notContainsQueries = append(notContainsQueries, "link NOT LIKE ?")
					notContainsArgs = append(notContainsArgs, "%"+v+"%")
				}
			}
			if len(notContainsQueries) > 0 {
				query = query.Where("("+strings.Join(notContainsQueries, " OR ")+")", notContainsArgs...)
			}
		}
		if f.Link.IsEmpty {
			query = query.Where("link = '' OR link IS NULL")
		}
		if f.Link.IsNotEmpty {
			query = query.Where("link <> '' AND link IS NOT NULL")
		}
	}
	if f.Pagination != nil {
		if f.Pagination.Page > 0 && f.Pagination.Limit > 0 {
			offset := (f.Pagination.Page - 1) * f.Pagination.Limit
			query = query.Limit(uint64(f.Pagination.Limit)).Offset(uint64(offset))
		} else if f.Pagination.Limit > 0 {
			query = query.Limit(uint64(f.Pagination.Limit))
		}
	}
	return query
}
func (f *FilterAddRequest) BuildMongoFilter() bson.M {
	if f == nil {
		return nil
	}
	filter := bson.M{}
	if f.Link != nil {
		fieldFilter := bson.M{}
		if f.Link.Eq != "" {
			fieldFilter["$eq"] = f.Link.Eq
		}
		if f.Link.Ne != "" {
			fieldFilter["$ne"] = f.Link.Ne
		}
		if f.Link.Lt != "" {
			fieldFilter["$lt"] = f.Link.Lt
		}
		if f.Link.Le != "" {
			fieldFilter["$lte"] = f.Link.Le
		}
		if f.Link.Gt != "" {
			fieldFilter["$gt"] = f.Link.Gt
		}
		if f.Link.Ge != "" {
			fieldFilter["$gte"] = f.Link.Ge
		}
		if len(f.Link.Contains) > 0 {
			fieldFilter["$in"] = f.Link.Contains
		}
		if len(f.Link.NotContains) > 0 {
			fieldFilter["$nin"] = f.Link.NotContains
		}
		if len(fieldFilter) > 0 {
			filter["link"] = fieldFilter
		}
	}
	return filter
}

type FilterAddResponse struct {
	Link       *StringFilterInput `json:"link"`
	Pagination *Pagination        `json:"pagination,omitempty"`
}

func (f *FilterAddResponse) BuildFilter(query squirrel.SelectBuilder) squirrel.SelectBuilder {
	if f.Link != nil {
		if f.Link.Eq != "" {
			query = query.Where("link = ?", f.Link.Eq)
		}
		if f.Link.Ne != "" {
			query = query.Where("link <> ?", f.Link.Ne)
		}
		if f.Link.Lt != "" {
			query = query.Where("link < ?", f.Link.Lt)
		}
		if f.Link.Le != "" {
			query = query.Where("link <= ?", f.Link.Le)
		}
		if f.Link.Gt != "" {
			query = query.Where("link > ?", f.Link.Gt)
		}
		if f.Link.Ge != "" {
			query = query.Where("link >= ?", f.Link.Ge)
		}
		if f.Link.StartsWith != "" {
			query = query.Where("link LIKE '%' || ?", f.Link.StartsWith)
		}
		if f.Link.EndsWith != "" {
			query = query.Where("link LIKE ? || '%'", f.Link.EndsWith)
		}
		if len(f.Link.Contains) > 0 {
			containsQueries := []string{}
			containsArgs := []interface{}{}
			for _, v := range f.Link.Contains {
				if v != "" {
					containsQueries = append(containsQueries, "link LIKE ?")
					containsArgs = append(containsArgs, "%"+v+"%")
				}
			}
			if len(containsQueries) > 0 {
				query = query.Where("("+strings.Join(containsQueries, " OR ")+")", containsArgs...)
			}
		}
		if len(f.Link.NotContains) > 0 {
			notContainsQueries := []string{}
			notContainsArgs := []interface{}{}
			for _, v := range f.Link.NotContains {
				if v != "" {
					notContainsQueries = append(notContainsQueries, "link NOT LIKE ?")
					notContainsArgs = append(notContainsArgs, "%"+v+"%")
				}
			}
			if len(notContainsQueries) > 0 {
				query = query.Where("("+strings.Join(notContainsQueries, " OR ")+")", notContainsArgs...)
			}
		}
		if f.Link.IsEmpty {
			query = query.Where("link = '' OR link IS NULL")
		}
		if f.Link.IsNotEmpty {
			query = query.Where("link <> '' AND link IS NOT NULL")
		}
	}
	if f.Pagination != nil {
		if f.Pagination.Page > 0 && f.Pagination.Limit > 0 {
			offset := (f.Pagination.Page - 1) * f.Pagination.Limit
			query = query.Limit(uint64(f.Pagination.Limit)).Offset(uint64(offset))
		} else if f.Pagination.Limit > 0 {
			query = query.Limit(uint64(f.Pagination.Limit))
		}
	}
	return query
}
func (f *FilterAddResponse) BuildMongoFilter() bson.M {
	if f == nil {
		return nil
	}
	filter := bson.M{}
	if f.Link != nil {
		fieldFilter := bson.M{}
		if f.Link.Eq != "" {
			fieldFilter["$eq"] = f.Link.Eq
		}
		if f.Link.Ne != "" {
			fieldFilter["$ne"] = f.Link.Ne
		}
		if f.Link.Lt != "" {
			fieldFilter["$lt"] = f.Link.Lt
		}
		if f.Link.Le != "" {
			fieldFilter["$lte"] = f.Link.Le
		}
		if f.Link.Gt != "" {
			fieldFilter["$gt"] = f.Link.Gt
		}
		if f.Link.Ge != "" {
			fieldFilter["$gte"] = f.Link.Ge
		}
		if len(f.Link.Contains) > 0 {
			fieldFilter["$in"] = f.Link.Contains
		}
		if len(f.Link.NotContains) > 0 {
			fieldFilter["$nin"] = f.Link.NotContains
		}
		if len(fieldFilter) > 0 {
			filter["link"] = fieldFilter
		}
	}
	return filter
}

type FilterUpdateRequest struct {
	Link       *StringFilterInput `json:"link"`
	Pagination *Pagination        `json:"pagination,omitempty"`
}

func (f *FilterUpdateRequest) BuildFilter(query squirrel.SelectBuilder) squirrel.SelectBuilder {
	if f.Link != nil {
		if f.Link.Eq != "" {
			query = query.Where("link = ?", f.Link.Eq)
		}
		if f.Link.Ne != "" {
			query = query.Where("link <> ?", f.Link.Ne)
		}
		if f.Link.Lt != "" {
			query = query.Where("link < ?", f.Link.Lt)
		}
		if f.Link.Le != "" {
			query = query.Where("link <= ?", f.Link.Le)
		}
		if f.Link.Gt != "" {
			query = query.Where("link > ?", f.Link.Gt)
		}
		if f.Link.Ge != "" {
			query = query.Where("link >= ?", f.Link.Ge)
		}
		if f.Link.StartsWith != "" {
			query = query.Where("link LIKE '%' || ?", f.Link.StartsWith)
		}
		if f.Link.EndsWith != "" {
			query = query.Where("link LIKE ? || '%'", f.Link.EndsWith)
		}
		if len(f.Link.Contains) > 0 {
			containsQueries := []string{}
			containsArgs := []interface{}{}
			for _, v := range f.Link.Contains {
				if v != "" {
					containsQueries = append(containsQueries, "link LIKE ?")
					containsArgs = append(containsArgs, "%"+v+"%")
				}
			}
			if len(containsQueries) > 0 {
				query = query.Where("("+strings.Join(containsQueries, " OR ")+")", containsArgs...)
			}
		}
		if len(f.Link.NotContains) > 0 {
			notContainsQueries := []string{}
			notContainsArgs := []interface{}{}
			for _, v := range f.Link.NotContains {
				if v != "" {
					notContainsQueries = append(notContainsQueries, "link NOT LIKE ?")
					notContainsArgs = append(notContainsArgs, "%"+v+"%")
				}
			}
			if len(notContainsQueries) > 0 {
				query = query.Where("("+strings.Join(notContainsQueries, " OR ")+")", notContainsArgs...)
			}
		}
		if f.Link.IsEmpty {
			query = query.Where("link = '' OR link IS NULL")
		}
		if f.Link.IsNotEmpty {
			query = query.Where("link <> '' AND link IS NOT NULL")
		}
	}
	if f.Pagination != nil {
		if f.Pagination.Page > 0 && f.Pagination.Limit > 0 {
			offset := (f.Pagination.Page - 1) * f.Pagination.Limit
			query = query.Limit(uint64(f.Pagination.Limit)).Offset(uint64(offset))
		} else if f.Pagination.Limit > 0 {
			query = query.Limit(uint64(f.Pagination.Limit))
		}
	}
	return query
}
func (f *FilterUpdateRequest) BuildMongoFilter() bson.M {
	if f == nil {
		return nil
	}
	filter := bson.M{}
	if f.Link != nil {
		fieldFilter := bson.M{}
		if f.Link.Eq != "" {
			fieldFilter["$eq"] = f.Link.Eq
		}
		if f.Link.Ne != "" {
			fieldFilter["$ne"] = f.Link.Ne
		}
		if f.Link.Lt != "" {
			fieldFilter["$lt"] = f.Link.Lt
		}
		if f.Link.Le != "" {
			fieldFilter["$lte"] = f.Link.Le
		}
		if f.Link.Gt != "" {
			fieldFilter["$gt"] = f.Link.Gt
		}
		if f.Link.Ge != "" {
			fieldFilter["$gte"] = f.Link.Ge
		}
		if len(f.Link.Contains) > 0 {
			fieldFilter["$in"] = f.Link.Contains
		}
		if len(f.Link.NotContains) > 0 {
			fieldFilter["$nin"] = f.Link.NotContains
		}
		if len(fieldFilter) > 0 {
			filter["link"] = fieldFilter
		}
	}
	return filter
}

type FilterUpdateResponse struct {
	Link       *StringFilterInput `json:"link"`
	Pagination *Pagination        `json:"pagination,omitempty"`
}

func (f *FilterUpdateResponse) BuildFilter(query squirrel.SelectBuilder) squirrel.SelectBuilder {
	if f.Link != nil {
		if f.Link.Eq != "" {
			query = query.Where("link = ?", f.Link.Eq)
		}
		if f.Link.Ne != "" {
			query = query.Where("link <> ?", f.Link.Ne)
		}
		if f.Link.Lt != "" {
			query = query.Where("link < ?", f.Link.Lt)
		}
		if f.Link.Le != "" {
			query = query.Where("link <= ?", f.Link.Le)
		}
		if f.Link.Gt != "" {
			query = query.Where("link > ?", f.Link.Gt)
		}
		if f.Link.Ge != "" {
			query = query.Where("link >= ?", f.Link.Ge)
		}
		if f.Link.StartsWith != "" {
			query = query.Where("link LIKE '%' || ?", f.Link.StartsWith)
		}
		if f.Link.EndsWith != "" {
			query = query.Where("link LIKE ? || '%'", f.Link.EndsWith)
		}
		if len(f.Link.Contains) > 0 {
			containsQueries := []string{}
			containsArgs := []interface{}{}
			for _, v := range f.Link.Contains {
				if v != "" {
					containsQueries = append(containsQueries, "link LIKE ?")
					containsArgs = append(containsArgs, "%"+v+"%")
				}
			}
			if len(containsQueries) > 0 {
				query = query.Where("("+strings.Join(containsQueries, " OR ")+")", containsArgs...)
			}
		}
		if len(f.Link.NotContains) > 0 {
			notContainsQueries := []string{}
			notContainsArgs := []interface{}{}
			for _, v := range f.Link.NotContains {
				if v != "" {
					notContainsQueries = append(notContainsQueries, "link NOT LIKE ?")
					notContainsArgs = append(notContainsArgs, "%"+v+"%")
				}
			}
			if len(notContainsQueries) > 0 {
				query = query.Where("("+strings.Join(notContainsQueries, " OR ")+")", notContainsArgs...)
			}
		}
		if f.Link.IsEmpty {
			query = query.Where("link = '' OR link IS NULL")
		}
		if f.Link.IsNotEmpty {
			query = query.Where("link <> '' AND link IS NOT NULL")
		}
	}
	if f.Pagination != nil {
		if f.Pagination.Page > 0 && f.Pagination.Limit > 0 {
			offset := (f.Pagination.Page - 1) * f.Pagination.Limit
			query = query.Limit(uint64(f.Pagination.Limit)).Offset(uint64(offset))
		} else if f.Pagination.Limit > 0 {
			query = query.Limit(uint64(f.Pagination.Limit))
		}
	}
	return query
}
func (f *FilterUpdateResponse) BuildMongoFilter() bson.M {
	if f == nil {
		return nil
	}
	filter := bson.M{}
	if f.Link != nil {
		fieldFilter := bson.M{}
		if f.Link.Eq != "" {
			fieldFilter["$eq"] = f.Link.Eq
		}
		if f.Link.Ne != "" {
			fieldFilter["$ne"] = f.Link.Ne
		}
		if f.Link.Lt != "" {
			fieldFilter["$lt"] = f.Link.Lt
		}
		if f.Link.Le != "" {
			fieldFilter["$lte"] = f.Link.Le
		}
		if f.Link.Gt != "" {
			fieldFilter["$gt"] = f.Link.Gt
		}
		if f.Link.Ge != "" {
			fieldFilter["$gte"] = f.Link.Ge
		}
		if len(f.Link.Contains) > 0 {
			fieldFilter["$in"] = f.Link.Contains
		}
		if len(f.Link.NotContains) > 0 {
			fieldFilter["$nin"] = f.Link.NotContains
		}
		if len(fieldFilter) > 0 {
			filter["link"] = fieldFilter
		}
	}
	return filter
}

type FilterDeleteRequest struct {
	Hash       *StringFilterInput `json:"hash"`
	Pagination *Pagination        `json:"pagination,omitempty"`
}

func (f *FilterDeleteRequest) BuildFilter(query squirrel.SelectBuilder) squirrel.SelectBuilder {
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
	if f.Pagination != nil {
		if f.Pagination.Page > 0 && f.Pagination.Limit > 0 {
			offset := (f.Pagination.Page - 1) * f.Pagination.Limit
			query = query.Limit(uint64(f.Pagination.Limit)).Offset(uint64(offset))
		} else if f.Pagination.Limit > 0 {
			query = query.Limit(uint64(f.Pagination.Limit))
		}
	}
	return query
}
func (f *FilterDeleteRequest) BuildMongoFilter() bson.M {
	if f == nil {
		return nil
	}
	filter := bson.M{}
	if f.Hash != nil {
		fieldFilter := bson.M{}
		if f.Hash.Eq != "" {
			fieldFilter["$eq"] = f.Hash.Eq
		}
		if f.Hash.Ne != "" {
			fieldFilter["$ne"] = f.Hash.Ne
		}
		if f.Hash.Lt != "" {
			fieldFilter["$lt"] = f.Hash.Lt
		}
		if f.Hash.Le != "" {
			fieldFilter["$lte"] = f.Hash.Le
		}
		if f.Hash.Gt != "" {
			fieldFilter["$gt"] = f.Hash.Gt
		}
		if f.Hash.Ge != "" {
			fieldFilter["$gte"] = f.Hash.Ge
		}
		if len(f.Hash.Contains) > 0 {
			fieldFilter["$in"] = f.Hash.Contains
		}
		if len(f.Hash.NotContains) > 0 {
			fieldFilter["$nin"] = f.Hash.NotContains
		}
		if len(fieldFilter) > 0 {
			filter["hash"] = fieldFilter
		}
	}
	return filter
}
