// Code generated by protoc-gen-go-orm. DO NOT EDIT.
// versions:
// - protoc-gen-go-orm v1.1.0
// - protoc             (unknown)
// source: infrastructure/rpc/link/v1/link.proto

package v1

import (
	"strings"
	"github.com/Masterminds/squirrel"
	"go.mongodb.org/mongo-driver/bson"
)

type FilterGetRequest struct {
	Hash       *StringFilterInput `json:"hash"`
	Pagination *Pagination        `json:"pagination,omitempty"`
}

func (f *FilterGetRequest) BuildFilter(query squirrel.SelectBuilder) squirrel.SelectBuilder {
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
func (f *FilterGetRequest) BuildMongoFilter() bson.M {
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

type FilterGetResponse struct {
	Link       *StringFilterInput `json:"link"`
	Pagination *Pagination        `json:"pagination,omitempty"`
}

func (f *FilterGetResponse) BuildFilter(query squirrel.SelectBuilder) squirrel.SelectBuilder {
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
func (f *FilterGetResponse) BuildMongoFilter() bson.M {
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

type FilterListRequest struct {
	Filter     *StringFilterInput `json:"filter"`
	Pagination *Pagination        `json:"pagination,omitempty"`
}

func (f *FilterListRequest) BuildFilter(query squirrel.SelectBuilder) squirrel.SelectBuilder {
	if f.Filter != nil {
		if f.Filter.Eq != "" {
			query = query.Where("filter = ?", f.Filter.Eq)
		}
		if f.Filter.Ne != "" {
			query = query.Where("filter <> ?", f.Filter.Ne)
		}
		if f.Filter.Lt != "" {
			query = query.Where("filter < ?", f.Filter.Lt)
		}
		if f.Filter.Le != "" {
			query = query.Where("filter <= ?", f.Filter.Le)
		}
		if f.Filter.Gt != "" {
			query = query.Where("filter > ?", f.Filter.Gt)
		}
		if f.Filter.Ge != "" {
			query = query.Where("filter >= ?", f.Filter.Ge)
		}
		if f.Filter.StartsWith != "" {
			query = query.Where("filter LIKE '%' || ?", f.Filter.StartsWith)
		}
		if f.Filter.EndsWith != "" {
			query = query.Where("filter LIKE ? || '%'", f.Filter.EndsWith)
		}
		if len(f.Filter.Contains) > 0 {
			containsQueries := []string{}
			containsArgs := []interface{}{}
			for _, v := range f.Filter.Contains {
				if v != "" {
					containsQueries = append(containsQueries, "filter LIKE ?")
					containsArgs = append(containsArgs, "%"+v+"%")
				}
			}
			if len(containsQueries) > 0 {
				query = query.Where("("+strings.Join(containsQueries, " OR ")+")", containsArgs...)
			}
		}
		if len(f.Filter.NotContains) > 0 {
			notContainsQueries := []string{}
			notContainsArgs := []interface{}{}
			for _, v := range f.Filter.NotContains {
				if v != "" {
					notContainsQueries = append(notContainsQueries, "filter NOT LIKE ?")
					notContainsArgs = append(notContainsArgs, "%"+v+"%")
				}
			}
			if len(notContainsQueries) > 0 {
				query = query.Where("("+strings.Join(notContainsQueries, " OR ")+")", notContainsArgs...)
			}
		}
		if f.Filter.IsEmpty {
			query = query.Where("filter = '' OR filter IS NULL")
		}
		if f.Filter.IsNotEmpty {
			query = query.Where("filter <> '' AND filter IS NOT NULL")
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
func (f *FilterListRequest) BuildMongoFilter() bson.M {
	if f == nil {
		return nil
	}
	filter := bson.M{}
	if f.Filter != nil {
		fieldFilter := bson.M{}
		if f.Filter.Eq != "" {
			fieldFilter["$eq"] = f.Filter.Eq
		}
		if f.Filter.Ne != "" {
			fieldFilter["$ne"] = f.Filter.Ne
		}
		if f.Filter.Lt != "" {
			fieldFilter["$lt"] = f.Filter.Lt
		}
		if f.Filter.Le != "" {
			fieldFilter["$lte"] = f.Filter.Le
		}
		if f.Filter.Gt != "" {
			fieldFilter["$gt"] = f.Filter.Gt
		}
		if f.Filter.Ge != "" {
			fieldFilter["$gte"] = f.Filter.Ge
		}
		if len(f.Filter.Contains) > 0 {
			fieldFilter["$in"] = f.Filter.Contains
		}
		if len(f.Filter.NotContains) > 0 {
			fieldFilter["$nin"] = f.Filter.NotContains
		}
		if len(fieldFilter) > 0 {
			filter["filter"] = fieldFilter
		}
	}
	return filter
}

type FilterListResponse struct {
	Links      *StringFilterInput `json:"links"`
	Pagination *Pagination        `json:"pagination,omitempty"`
}

func (f *FilterListResponse) BuildFilter(query squirrel.SelectBuilder) squirrel.SelectBuilder {
	if f.Links != nil {
		if f.Links.Eq != "" {
			query = query.Where("links = ?", f.Links.Eq)
		}
		if f.Links.Ne != "" {
			query = query.Where("links <> ?", f.Links.Ne)
		}
		if f.Links.Lt != "" {
			query = query.Where("links < ?", f.Links.Lt)
		}
		if f.Links.Le != "" {
			query = query.Where("links <= ?", f.Links.Le)
		}
		if f.Links.Gt != "" {
			query = query.Where("links > ?", f.Links.Gt)
		}
		if f.Links.Ge != "" {
			query = query.Where("links >= ?", f.Links.Ge)
		}
		if f.Links.StartsWith != "" {
			query = query.Where("links LIKE '%' || ?", f.Links.StartsWith)
		}
		if f.Links.EndsWith != "" {
			query = query.Where("links LIKE ? || '%'", f.Links.EndsWith)
		}
		if len(f.Links.Contains) > 0 {
			containsQueries := []string{}
			containsArgs := []interface{}{}
			for _, v := range f.Links.Contains {
				if v != "" {
					containsQueries = append(containsQueries, "links LIKE ?")
					containsArgs = append(containsArgs, "%"+v+"%")
				}
			}
			if len(containsQueries) > 0 {
				query = query.Where("("+strings.Join(containsQueries, " OR ")+")", containsArgs...)
			}
		}
		if len(f.Links.NotContains) > 0 {
			notContainsQueries := []string{}
			notContainsArgs := []interface{}{}
			for _, v := range f.Links.NotContains {
				if v != "" {
					notContainsQueries = append(notContainsQueries, "links NOT LIKE ?")
					notContainsArgs = append(notContainsArgs, "%"+v+"%")
				}
			}
			if len(notContainsQueries) > 0 {
				query = query.Where("("+strings.Join(notContainsQueries, " OR ")+")", notContainsArgs...)
			}
		}
		if f.Links.IsEmpty {
			query = query.Where("links = '' OR links IS NULL")
		}
		if f.Links.IsNotEmpty {
			query = query.Where("links <> '' AND links IS NOT NULL")
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
func (f *FilterListResponse) BuildMongoFilter() bson.M {
	if f == nil {
		return nil
	}
	filter := bson.M{}
	if f.Links != nil {
		fieldFilter := bson.M{}
		if f.Links.Eq != "" {
			fieldFilter["$eq"] = f.Links.Eq
		}
		if f.Links.Ne != "" {
			fieldFilter["$ne"] = f.Links.Ne
		}
		if f.Links.Lt != "" {
			fieldFilter["$lt"] = f.Links.Lt
		}
		if f.Links.Le != "" {
			fieldFilter["$lte"] = f.Links.Le
		}
		if f.Links.Gt != "" {
			fieldFilter["$gt"] = f.Links.Gt
		}
		if f.Links.Ge != "" {
			fieldFilter["$gte"] = f.Links.Ge
		}
		if len(f.Links.Contains) > 0 {
			fieldFilter["$in"] = f.Links.Contains
		}
		if len(f.Links.NotContains) > 0 {
			fieldFilter["$nin"] = f.Links.NotContains
		}
		if len(fieldFilter) > 0 {
			filter["links"] = fieldFilter
		}
	}
	return filter
}

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
