package util

import (
	"fmt"
	"reflect"
)

func SetField(v any, fieldName string, fieldValue any) {
	val := reflect.ValueOf(v)
	if val.Type().Kind() == reflect.Pointer {
		val = val.Elem()
	}
	if val.Type().Kind() != reflect.Struct {
		return
	}
	f := val.FieldByName(fieldName)
	if !f.IsValid() {
		panic(fmt.Errorf("no such field: %s in obj", fieldName))
	}
	if !f.CanSet() {
		panic(fmt.Errorf("cannot set field %s", fieldName))
	}
	// Ensure the value type matches the field type
	if f.Type() != val.Type() {
		panic(fmt.Errorf("provided value type didn't match obj field type"))
	}
	f.Set(reflect.ValueOf(fieldValue))
}
