package util

import (
	"reflect"
)

func IsStruct(v any) bool {
	t := reflect.TypeOf(v)
	switch t.Kind() {
	case reflect.Struct:
		return true
	case reflect.Pointer:
		t = t.Elem()
		return t.Kind() == reflect.Struct
	default:
		return false
	}
}
