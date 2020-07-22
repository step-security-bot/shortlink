// Code generated by protoc-gen-gotemplate. DO NOT EDIT.
// source: pkg/link/link.proto

package query

type Filter struct { // nolint unused
	Url        *StringFilterInput
	Hash       *StringFilterInput
	Describe   *StringFilterInput
	Created_at *StringFilterInput
	Updated_at *StringFilterInput
	Link       *StringFilterInput
}

type StringFilterInput struct { // nolint unused
	Eq          *string
	Ne          *string
	Lt          *string
	Le          *string
	Gt          *string
	Ge          *string
	Contains    *string
	NotContains *string
}

// GetKeys - return all keys
func (f *Filter) GetKeys() []string {
	return []string{
		"Url",
		"Hash",
		"Describe",
		"Created_at",
		"Updated_at",
		"Link",
	}
}
