package util

import (
	"reflect"
)

func IsArray(v any) bool {
	t := reflect.TypeOf(v)
	switch t.Kind() {
	case reflect.Slice:
		return true
	case reflect.Array:
		return true
	default:
		return false
	}
}
