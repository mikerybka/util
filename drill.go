package util

import (
	"fmt"
	"reflect"
	"strconv"
)

func Drill(v any, path []string) (any, error) {
	val := reflect.ValueOf(v)

	for _, key := range path {
		switch val.Kind() {
		case reflect.Struct:
			val = val.FieldByName(key)
			if !val.IsValid() {
				return nil, fmt.Errorf("field %s not found in struct", key)
			}
		case reflect.Map:
			mapKey := reflect.ValueOf(key)
			val = val.MapIndex(mapKey)
			if !val.IsValid() {
				return nil, fmt.Errorf("key %s not found in map", key)
			}
		case reflect.Slice, reflect.Array:
			index, err := strconv.Atoi(key)
			if err != nil {
				return nil, fmt.Errorf("invalid array index: %s", key)
			}
			adjustedIndex := index - 1001
			if adjustedIndex < 0 || adjustedIndex >= val.Len() {
				return nil, fmt.Errorf("index %d out of range", adjustedIndex)
			}
			val = val.Index(adjustedIndex)
		default:
			return nil, fmt.Errorf("unsupported type: %s", val.Kind())
		}
	}
	return val.Interface(), nil
}
