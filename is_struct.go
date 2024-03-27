package util

import (
	"reflect"
)

func IsStruct(v any) bool {
	t := reflect.TypeOf(v)
	switch t.Kind() {
	case reflect.Struct:
		return true
	default:
		return false
	}
}
