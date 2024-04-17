// Code generated by protoc-gen-go-orm. DO NOT EDIT.
// versions:
// - protoc-gen-go-orm v1.6.0
// - protoc             (unknown)
// source: domain/link/v1/link.proto

package filter

import (
	"reflect"
	"strings"

	domain "github.com/shortlink-org/shortlink/boundaries/link/link/infrastructure/repository/crud/types"
)

func (f *FilterLink) BuildRAMFilter(item any) bool {
	var fieldVal *domain.StringFilterInput
	var ok bool
	var v reflect.Value

	fieldVal, ok = reflect.ValueOf(f).Elem().FieldByName("FieldMask").Interface().(*domain.StringFilterInput)
	if !ok || fieldVal == nil {
		return true
	} // If field is not found or nil, no filtering is applied
	v = reflect.ValueOf(item).Elem().FieldByName("FieldMask")
	if !v.IsValid() {
		return false
	} // If the link does not have this field, fail the filter
	if fieldVal.Eq != "" && v.String() != fieldVal.Eq {
		return false
	}
	if fieldVal.Ne != "" && v.String() == fieldVal.Ne {
		return false
	}
	if fieldVal.Lt != "" && !(v.String() < fieldVal.Lt) {
		return false
	}
	if fieldVal.Le != "" && !(v.String() <= fieldVal.Le) {
		return false
	}
	if fieldVal.Gt != "" && !(v.String() > fieldVal.Gt) {
		return false
	}
	if fieldVal.Ge != "" && !(v.String() >= fieldVal.Ge) {
		return false
	}
	if fieldVal.StartsWith != "" && !strings.HasPrefix(v.String(), fieldVal.StartsWith) {
		return false
	}
	if fieldVal.EndsWith != "" && !strings.HasSuffix(v.String(), fieldVal.EndsWith) {
		return false
	}
	for _, contain := range fieldVal.Contains {
		if !strings.Contains(v.String(), contain) {
			return false
		}
	}
	for _, notContain := range fieldVal.NotContains {
		if strings.Contains(v.String(), notContain) {
			return false
		}
	}

	if fieldVal.IsEmpty && v.String() != "" {
		return false
	}
	if fieldVal.IsNotEmpty && v.String() == "" {
		return false
	}
	fieldVal, ok = reflect.ValueOf(f).Elem().FieldByName("Url").Interface().(*domain.StringFilterInput)
	if !ok || fieldVal == nil {
		return true
	} // If field is not found or nil, no filtering is applied
	v = reflect.ValueOf(item).Elem().FieldByName("Url")
	if !v.IsValid() {
		return false
	} // If the link does not have this field, fail the filter
	if fieldVal.Eq != "" && v.String() != fieldVal.Eq {
		return false
	}
	if fieldVal.Ne != "" && v.String() == fieldVal.Ne {
		return false
	}
	if fieldVal.Lt != "" && !(v.String() < fieldVal.Lt) {
		return false
	}
	if fieldVal.Le != "" && !(v.String() <= fieldVal.Le) {
		return false
	}
	if fieldVal.Gt != "" && !(v.String() > fieldVal.Gt) {
		return false
	}
	if fieldVal.Ge != "" && !(v.String() >= fieldVal.Ge) {
		return false
	}
	if fieldVal.StartsWith != "" && !strings.HasPrefix(v.String(), fieldVal.StartsWith) {
		return false
	}
	if fieldVal.EndsWith != "" && !strings.HasSuffix(v.String(), fieldVal.EndsWith) {
		return false
	}
	for _, contain := range fieldVal.Contains {
		if !strings.Contains(v.String(), contain) {
			return false
		}
	}
	for _, notContain := range fieldVal.NotContains {
		if strings.Contains(v.String(), notContain) {
			return false
		}
	}

	if fieldVal.IsEmpty && v.String() != "" {
		return false
	}
	if fieldVal.IsNotEmpty && v.String() == "" {
		return false
	}
	fieldVal, ok = reflect.ValueOf(f).Elem().FieldByName("Hash").Interface().(*domain.StringFilterInput)
	if !ok || fieldVal == nil {
		return true
	} // If field is not found or nil, no filtering is applied
	v = reflect.ValueOf(item).Elem().FieldByName("Hash")
	if !v.IsValid() {
		return false
	} // If the link does not have this field, fail the filter
	if fieldVal.Eq != "" && v.String() != fieldVal.Eq {
		return false
	}
	if fieldVal.Ne != "" && v.String() == fieldVal.Ne {
		return false
	}
	if fieldVal.Lt != "" && !(v.String() < fieldVal.Lt) {
		return false
	}
	if fieldVal.Le != "" && !(v.String() <= fieldVal.Le) {
		return false
	}
	if fieldVal.Gt != "" && !(v.String() > fieldVal.Gt) {
		return false
	}
	if fieldVal.Ge != "" && !(v.String() >= fieldVal.Ge) {
		return false
	}
	if fieldVal.StartsWith != "" && !strings.HasPrefix(v.String(), fieldVal.StartsWith) {
		return false
	}
	if fieldVal.EndsWith != "" && !strings.HasSuffix(v.String(), fieldVal.EndsWith) {
		return false
	}
	for _, contain := range fieldVal.Contains {
		if !strings.Contains(v.String(), contain) {
			return false
		}
	}
	for _, notContain := range fieldVal.NotContains {
		if strings.Contains(v.String(), notContain) {
			return false
		}
	}

	if fieldVal.IsEmpty && v.String() != "" {
		return false
	}
	if fieldVal.IsNotEmpty && v.String() == "" {
		return false
	}
	fieldVal, ok = reflect.ValueOf(f).Elem().FieldByName("Describe").Interface().(*domain.StringFilterInput)
	if !ok || fieldVal == nil {
		return true
	} // If field is not found or nil, no filtering is applied
	v = reflect.ValueOf(item).Elem().FieldByName("Describe")
	if !v.IsValid() {
		return false
	} // If the link does not have this field, fail the filter
	if fieldVal.Eq != "" && v.String() != fieldVal.Eq {
		return false
	}
	if fieldVal.Ne != "" && v.String() == fieldVal.Ne {
		return false
	}
	if fieldVal.Lt != "" && !(v.String() < fieldVal.Lt) {
		return false
	}
	if fieldVal.Le != "" && !(v.String() <= fieldVal.Le) {
		return false
	}
	if fieldVal.Gt != "" && !(v.String() > fieldVal.Gt) {
		return false
	}
	if fieldVal.Ge != "" && !(v.String() >= fieldVal.Ge) {
		return false
	}
	if fieldVal.StartsWith != "" && !strings.HasPrefix(v.String(), fieldVal.StartsWith) {
		return false
	}
	if fieldVal.EndsWith != "" && !strings.HasSuffix(v.String(), fieldVal.EndsWith) {
		return false
	}
	for _, contain := range fieldVal.Contains {
		if !strings.Contains(v.String(), contain) {
			return false
		}
	}
	for _, notContain := range fieldVal.NotContains {
		if strings.Contains(v.String(), notContain) {
			return false
		}
	}

	if fieldVal.IsEmpty && v.String() != "" {
		return false
	}
	if fieldVal.IsNotEmpty && v.String() == "" {
		return false
	}
	fieldVal, ok = reflect.ValueOf(f).Elem().FieldByName("CreatedAt").Interface().(*domain.StringFilterInput)
	if !ok || fieldVal == nil {
		return true
	} // If field is not found or nil, no filtering is applied
	v = reflect.ValueOf(item).Elem().FieldByName("CreatedAt")
	if !v.IsValid() {
		return false
	} // If the link does not have this field, fail the filter
	if fieldVal.Eq != "" && v.String() != fieldVal.Eq {
		return false
	}
	if fieldVal.Ne != "" && v.String() == fieldVal.Ne {
		return false
	}
	if fieldVal.Lt != "" && !(v.String() < fieldVal.Lt) {
		return false
	}
	if fieldVal.Le != "" && !(v.String() <= fieldVal.Le) {
		return false
	}
	if fieldVal.Gt != "" && !(v.String() > fieldVal.Gt) {
		return false
	}
	if fieldVal.Ge != "" && !(v.String() >= fieldVal.Ge) {
		return false
	}
	if fieldVal.StartsWith != "" && !strings.HasPrefix(v.String(), fieldVal.StartsWith) {
		return false
	}
	if fieldVal.EndsWith != "" && !strings.HasSuffix(v.String(), fieldVal.EndsWith) {
		return false
	}
	for _, contain := range fieldVal.Contains {
		if !strings.Contains(v.String(), contain) {
			return false
		}
	}
	for _, notContain := range fieldVal.NotContains {
		if strings.Contains(v.String(), notContain) {
			return false
		}
	}

	if fieldVal.IsEmpty && v.String() != "" {
		return false
	}
	if fieldVal.IsNotEmpty && v.String() == "" {
		return false
	}
	fieldVal, ok = reflect.ValueOf(f).Elem().FieldByName("UpdatedAt").Interface().(*domain.StringFilterInput)
	if !ok || fieldVal == nil {
		return true
	} // If field is not found or nil, no filtering is applied
	v = reflect.ValueOf(item).Elem().FieldByName("UpdatedAt")
	if !v.IsValid() {
		return false
	} // If the link does not have this field, fail the filter
	if fieldVal.Eq != "" && v.String() != fieldVal.Eq {
		return false
	}
	if fieldVal.Ne != "" && v.String() == fieldVal.Ne {
		return false
	}
	if fieldVal.Lt != "" && !(v.String() < fieldVal.Lt) {
		return false
	}
	if fieldVal.Le != "" && !(v.String() <= fieldVal.Le) {
		return false
	}
	if fieldVal.Gt != "" && !(v.String() > fieldVal.Gt) {
		return false
	}
	if fieldVal.Ge != "" && !(v.String() >= fieldVal.Ge) {
		return false
	}
	if fieldVal.StartsWith != "" && !strings.HasPrefix(v.String(), fieldVal.StartsWith) {
		return false
	}
	if fieldVal.EndsWith != "" && !strings.HasSuffix(v.String(), fieldVal.EndsWith) {
		return false
	}
	for _, contain := range fieldVal.Contains {
		if !strings.Contains(v.String(), contain) {
			return false
		}
	}
	for _, notContain := range fieldVal.NotContains {
		if strings.Contains(v.String(), notContain) {
			return false
		}
	}

	if fieldVal.IsEmpty && v.String() != "" {
		return false
	}
	if fieldVal.IsNotEmpty && v.String() == "" {
		return false
	}
	return true
}
