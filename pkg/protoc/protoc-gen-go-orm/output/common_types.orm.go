// Code generated by protoc-gen-go-orm. DO NOT EDIT.
// versions:
// - protoc-gen-go-orm v1.6.0
// - protoc             v5.26.1
// source: fixtures/link.proto

package example

import (
	domain "github.com/shortlink-org/shortlink/pkg/protoc/protoc-gen-rich-model/fixtures"
)

type FilterLink domain.FilterLink

func NewFilter(params *types.FilterLink) *FilterLink {
	if params == nil {
		return nil
	}

	return &FilterLink{
		FieldMask: params.FieldMask,
		Url:       params.Url,
		Hash:      params.Hash,
		Describe:  params.Describe,
		CreatedAt: params.CreatedAt,
		UpdatedAt: params.UpdatedAt,
	}
}
