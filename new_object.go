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
		return &Array{
			Path:  path,
			Value: v.([]any),
		}
	}
	if IsMap(v) {
		return &Map{
			Path:  path,
			Value: v.(map[string]any),
		}
	}
	if IsStruct(v) {
		return &Struct{
			Path:  path,
			Value: StructToMap(v),
		}
	}
	panic(fmt.Errorf("unknown type for %s: %v", JoinPath(path), v))
}
