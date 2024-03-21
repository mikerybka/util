package util

import (
	"reflect"
)

func IsMap(v any) bool {
	t := reflect.TypeOf(v)
	switch t.Kind() {
	case reflect.Map:
		return true
	default:
		return false
	}
}
