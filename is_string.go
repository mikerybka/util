package util

import "reflect"

func IsString(v any) bool {
	t := reflect.TypeOf(v)
	switch t.Kind() {
	case reflect.String:
		return true
	default:
		return false
	}
}
