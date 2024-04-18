// Code generated by protoc-gen-go-orm. DO NOT EDIT.
// versions:
// - protoc-gen-go-orm v1.6.0
// - protoc             (unknown)
// source: domain/link/v1/link.proto

package filter

import (
	"go.mongodb.org/mongo-driver/bson"
)

func (f *FilterLink) BuildMongoFilter() bson.M {
	if f == nil {
		return nil
	}
	filter := bson.M{}
	if f.FieldMask != nil {
		fieldFilter := bson.M{}
		if f.FieldMask.Eq != "" {
			fieldFilter["$eq"] = f.FieldMask.Eq
		}
		if f.FieldMask.Ne != "" {
			fieldFilter["$ne"] = f.FieldMask.Ne
		}
		if f.FieldMask.Lt != "" {
			fieldFilter["$lt"] = f.FieldMask.Lt
		}
		if f.FieldMask.Le != "" {
			fieldFilter["$lte"] = f.FieldMask.Le
		}
		if f.FieldMask.Gt != "" {
			fieldFilter["$gt"] = f.FieldMask.Gt
		}
		if f.FieldMask.Ge != "" {
			fieldFilter["$gte"] = f.FieldMask.Ge
		}
		if len(f.FieldMask.Contains) > 0 {
			fieldFilter["$in"] = f.FieldMask.Contains
		}
		if len(f.FieldMask.NotContains) > 0 {
			fieldFilter["$nin"] = f.FieldMask.NotContains
		}
		if len(fieldFilter) > 0 {
			filter["fieldmask"] = fieldFilter
		}
	}
	if f.Url != nil {
		fieldFilter := bson.M{}
		if f.Url.Eq != "" {
			fieldFilter["$eq"] = f.Url.Eq
		}
		if f.Url.Ne != "" {
			fieldFilter["$ne"] = f.Url.Ne
		}
		if f.Url.Lt != "" {
			fieldFilter["$lt"] = f.Url.Lt
		}
		if f.Url.Le != "" {
			fieldFilter["$lte"] = f.Url.Le
		}
		if f.Url.Gt != "" {
			fieldFilter["$gt"] = f.Url.Gt
		}
		if f.Url.Ge != "" {
			fieldFilter["$gte"] = f.Url.Ge
		}
		if len(f.Url.Contains) > 0 {
			fieldFilter["$in"] = f.Url.Contains
		}
		if len(f.Url.NotContains) > 0 {
			fieldFilter["$nin"] = f.Url.NotContains
		}
		if len(fieldFilter) > 0 {
			filter["url"] = fieldFilter
		}
	}
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
	if f.Describe != nil {
		fieldFilter := bson.M{}
		if f.Describe.Eq != "" {
			fieldFilter["$eq"] = f.Describe.Eq
		}
		if f.Describe.Ne != "" {
			fieldFilter["$ne"] = f.Describe.Ne
		}
		if f.Describe.Lt != "" {
			fieldFilter["$lt"] = f.Describe.Lt
		}
		if f.Describe.Le != "" {
			fieldFilter["$lte"] = f.Describe.Le
		}
		if f.Describe.Gt != "" {
			fieldFilter["$gt"] = f.Describe.Gt
		}
		if f.Describe.Ge != "" {
			fieldFilter["$gte"] = f.Describe.Ge
		}
		if len(f.Describe.Contains) > 0 {
			fieldFilter["$in"] = f.Describe.Contains
		}
		if len(f.Describe.NotContains) > 0 {
			fieldFilter["$nin"] = f.Describe.NotContains
		}
		if len(fieldFilter) > 0 {
			filter["describe"] = fieldFilter
		}
	}
	if f.CreatedAt != nil {
		fieldFilter := bson.M{}
		if f.CreatedAt.Eq != "" {
			fieldFilter["$eq"] = f.CreatedAt.Eq
		}
		if f.CreatedAt.Ne != "" {
			fieldFilter["$ne"] = f.CreatedAt.Ne
		}
		if f.CreatedAt.Lt != "" {
			fieldFilter["$lt"] = f.CreatedAt.Lt
		}
		if f.CreatedAt.Le != "" {
			fieldFilter["$lte"] = f.CreatedAt.Le
		}
		if f.CreatedAt.Gt != "" {
			fieldFilter["$gt"] = f.CreatedAt.Gt
		}
		if f.CreatedAt.Ge != "" {
			fieldFilter["$gte"] = f.CreatedAt.Ge
		}
		if len(f.CreatedAt.Contains) > 0 {
			fieldFilter["$in"] = f.CreatedAt.Contains
		}
		if len(f.CreatedAt.NotContains) > 0 {
			fieldFilter["$nin"] = f.CreatedAt.NotContains
		}
		if len(fieldFilter) > 0 {
			filter["createdat"] = fieldFilter
		}
	}
	if f.UpdatedAt != nil {
		fieldFilter := bson.M{}
		if f.UpdatedAt.Eq != "" {
			fieldFilter["$eq"] = f.UpdatedAt.Eq
		}
		if f.UpdatedAt.Ne != "" {
			fieldFilter["$ne"] = f.UpdatedAt.Ne
		}
		if f.UpdatedAt.Lt != "" {
			fieldFilter["$lt"] = f.UpdatedAt.Lt
		}
		if f.UpdatedAt.Le != "" {
			fieldFilter["$lte"] = f.UpdatedAt.Le
		}
		if f.UpdatedAt.Gt != "" {
			fieldFilter["$gt"] = f.UpdatedAt.Gt
		}
		if f.UpdatedAt.Ge != "" {
			fieldFilter["$gte"] = f.UpdatedAt.Ge
		}
		if len(f.UpdatedAt.Contains) > 0 {
			fieldFilter["$in"] = f.UpdatedAt.Contains
		}
		if len(f.UpdatedAt.NotContains) > 0 {
			fieldFilter["$nin"] = f.UpdatedAt.NotContains
		}
		if len(fieldFilter) > 0 {
			filter["updatedat"] = fieldFilter
		}
	}
	return filter
}