package util

import "fmt"

func NewObject(path []string, v any) Object {
	if IsInt(v) {
		return &Int{
			Path:  path,
			Value: v.(int64),
		}
	}
	if IsString(v) {
		return &String{
			Path:  path,
			Value: v.(string),
		}
	}
	if IsBool(v) {
		return &Bool{
			Path:  path,
			Value: v.(bool),
		}
	}
	if IsArray(v) {
	}
	if IsMap(v) {
	}
	if IsStruct(v) {
	}
	panic(fmt.Errorf("unknown type for %s: %v", JoinPath(path), v))
}
