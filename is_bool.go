package util

import "reflect"

func IsBool(v any) bool {
	t := reflect.TypeOf(v)
	switch t.Kind() {
	case reflect.Bool:
		return true
	default:
		return false
	}
}
