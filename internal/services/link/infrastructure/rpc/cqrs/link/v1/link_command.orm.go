// Code generated by protoc-gen-go-orm. DO NOT EDIT.
// versions:
// - protoc-gen-go-orm v1.0.0
// - protoc             (unknown)
// source: infrastructure/rpc/cqrs/link/v1/link_command.proto

package v1

type FilterAddRequest struct {
	Link       *StringFilterInput `json:"link"`
	Pagination *Pagination        `json:"pagination,omitempty"`
}

type FilterAddResponse struct {
	Link       *StringFilterInput `json:"link"`
	Pagination *Pagination        `json:"pagination,omitempty"`
}

type FilterUpdateRequest struct {
	Link       *StringFilterInput `json:"link"`
	Pagination *Pagination        `json:"pagination,omitempty"`
}

type FilterUpdateResponse struct {
	Link       *StringFilterInput `json:"link"`
	Pagination *Pagination        `json:"pagination,omitempty"`
}

type FilterDeleteRequest struct {
	Hash       *StringFilterInput `json:"hash"`
	Pagination *Pagination        `json:"pagination,omitempty"`
}
