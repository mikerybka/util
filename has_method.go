package util

import "reflect"

func HasMethod(v any, m string) bool {
	t := reflect.TypeOf(v)
	_, ok := t.MethodByName(m)
	return ok
}
