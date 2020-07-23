package ram

import (
	"reflect"
	"strings"

	"github.com/batazor/shortlink/internal/store/query"
	"github.com/batazor/shortlink/pkg/link"
)

func isFilterSuccess(link *link.Link, filter *query.Filter) bool {
	// Skip empty filter
	if filter == nil {
		return true
	}

	r := reflect.ValueOf(filter)
	l := reflect.ValueOf(link)

	for _, key := range filter.GetKeys() {
		val := reflect.Indirect(r).FieldByName(key).Interface().(*query.StringFilterInput)

		// Skip empty value
		if val == nil {
			continue
		}

		// Eq
		if val.Eq != nil {
			linkValue := reflect.Indirect(l).FieldByName(key).Interface().(string)
			if linkValue != *val.Eq {
				return false
			}
		}

		// Ne
		if val.Ne != nil {
			linkValue := reflect.Indirect(l).FieldByName(key).Interface().(string)
			if linkValue == *val.Ne {
				return false
			}
		}

		// Lt
		if val.Lt != nil {
			linkValue := reflect.Indirect(l).FieldByName(key).Interface().(string)
			if linkValue > *val.Lt {
				return false
			}
		}

		// Le
		if val.Le != nil {
			linkValue := reflect.Indirect(l).FieldByName(key).Interface().(string)
			if linkValue >= *val.Le {
				return false
			}
		}

		// Gt
		if val.Gt != nil {
			linkValue := reflect.Indirect(l).FieldByName(key).Interface().(string)
			if linkValue < *val.Gt {
				return false
			}
		}

		// Ge
		if val.Ge != nil {
			linkValue := reflect.Indirect(l).FieldByName(key).Interface().(string)
			if linkValue <= *val.Ge {
				return false
			}
		}

		// Contains
		if val.Contains != nil {
			linkValue := reflect.Indirect(l).FieldByName(key).Interface().(string)
			if strings.Index(linkValue, *val.Contains) == -1 {
				return false
			}
		}

		// NotContains
		if val.Contains != nil {
			linkValue := reflect.Indirect(l).FieldByName(key).Interface().(string)
			if strings.Index(linkValue, *val.NotContains) != -1 {
				return false
			}
		}
	}

	return true
}
