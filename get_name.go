package util

import "reflect"

func GetName(v any) string {
	val := reflect.ValueOf(v)
	val.Kind()
}
