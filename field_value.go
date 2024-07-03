package util

import "reflect"

func FieldValue(v any, f string) any {
	rv := reflect.ValueOf(v)
	if rv.Type().Kind() == reflect.Pointer {
		rv = rv.Elem()
	}
	return rv.FieldByName(f).Interface()
}
