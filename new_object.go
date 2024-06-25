package util

import "fmt"

func NewObject(path []string, v any) Object {
	if v == nil {
		return &Nil{
			Path: path,
		}
	}

	i, ok := v.(int64)
	if ok {
		return &Int{
			Path:  path,
			Value: i,
		}
	}

	s, ok := v.(string)
	if ok {
		return &String{
			Path:  path,
			Value: s,
		}
	}

	b, ok := v.(bool)
	if ok {
		return &Bool{
			Path:  path,
			Value: b,
		}
	}

	if IsArray(v) {
		return &Array{
			Path:  path,
			Value: v,
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
