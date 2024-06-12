package util

import "reflect"

func GetName(v any) string {
	val := reflect.ValueOf(v)
	if val.Kind() == reflect.Pointer {
		return GetName(val.Elem().Interface())
	}
	if val.Kind() == reflect.Struct {
		name := val.FieldByName("Name")
		if !name.IsZero() {
			return name.String()
		}
		name = val.FieldByName("Title")
		if !name.IsZero() {
			return name.String()
		}
		name = val.FieldByName("ID")
		if !name.IsZero() {
			return name.String()
		}
	}
	return val.Type().String()
}
