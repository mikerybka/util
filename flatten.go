package util

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
)

// Flatten any Go object into a map[string]string.
// Paths are '/'-separated.
func Flatten(v any) map[string]string {
	result := make(map[string]string)
	flatten("", reflect.ValueOf(v), result)
	return result
}

func flatten(prefix string, v reflect.Value, result map[string]string) {
	switch v.Kind() {
	case reflect.Bool:
		result[prefix] = strconv.FormatBool(v.Bool())
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		result[prefix] = strconv.FormatInt(v.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		result[prefix] = strconv.FormatUint(v.Uint(), 10)
	case reflect.Float32, reflect.Float64:
		result[prefix] = strconv.FormatFloat(v.Float(), 'f', -1, 64)
	case reflect.String:
		result[prefix] = v.String()
	case reflect.Array, reflect.Slice:
		for i := 0; i < v.Len(); i++ {
			flatten(fmt.Sprintf("%s/%d", prefix, i), v.Index(i), result)
		}
	case reflect.Map:
		for _, key := range v.MapKeys() {
			flatten(fmt.Sprintf("%s/%s", prefix, key), v.MapIndex(key), result)
		}
	case reflect.Struct:
		t := v.Type()
		for i := 0; i < v.NumField(); i++ {
			field := t.Field(i)
			if field.PkgPath == "" { // Exported field
				flatten(fmt.Sprintf("%s/%s", prefix, field.Name), v.Field(i), result)
			}
		}
	default:
		// Convert non-scalar values to JSON strings
		jsonBytes, err := json.Marshal(v.Interface())
		if err != nil {
			return
		}
		result[prefix] = string(jsonBytes)
	}
}
