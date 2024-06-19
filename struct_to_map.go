package util

import "reflect"

func StructToMap(val any) map[string]any {
	m := map[string]any{}
	v := reflect.ValueOf(val)
	t := v.Type()
	if t.Kind() == reflect.Pointer {
		return StructToMap(v.Elem().Interface())
	}
	for i := 0; i < t.NumField(); i++ {
		m[t.Field(i).Name] = v.Field(i).Interface()
	}
	return m
}
