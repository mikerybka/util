package util

import "reflect"

func AppendUsingReflect(slice any, value any) any {
	sliceValue := reflect.ValueOf(slice)
	newSlice := reflect.MakeSlice(sliceValue.Type(), sliceValue.Len()+1, sliceValue.Cap()+1)
	reflect.Copy(newSlice, sliceValue)
	newValue := reflect.ValueOf(value)
	newSlice.Index(newSlice.Len() - 1).Set(newValue)
	return newSlice.Interface()
}
