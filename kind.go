package util

import "reflect"

// Kind determines the kind of a given value.
// It returns one of "struct", "map", "list", "scalar" or "null".
func Kind(v interface{}) string {
	if v == nil {
		return "null"
	}
	val := reflect.ValueOf(v)
	switch val.Kind() {
	case reflect.Struct:
		return "struct"
	case reflect.Map:
		return "map"
	case reflect.Array, reflect.Slice:
		return "list"
	case reflect.Int, reflect.Float32, reflect.Float64, reflect.Bool, reflect.String, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return "scalar"
	default:
		return "unknown"
	}
}
