package util

import (
	"reflect"
	"strconv"
)

func Drill(v any, part string) any {
	val := reflect.ValueOf(v)
	switch val.Type().Kind() {
	case reflect.Array, reflect.Slice:
		i, err := strconv.ParseInt(part, 10, 0)
		if err != nil {
			panic(err)
		}
		return val.Index(int(i)).Interface()
	case reflect.Map:
		switch val.Type().Key().Kind() {
		case reflect.String:
			return val.MapIndex(reflect.ValueOf(part)).Interface()
		default:
			panic("map keys must be strings")
		}
	case reflect.Struct:
		return val.FieldByName(part).Interface()
	default:
		panic("cannot drill")
	}
}
