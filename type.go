package util

import (
	"io"
	"reflect"
)

func GetType(v any) *Type {
	t := reflect.TypeOf(v)
	switch t.Kind() {
	case reflect.Invalid:
		panic("invalid type")
	case reflect.Array:
		return &Type{
			IsArray:  true,
			ElemType: GetType(reflect.New(t.Elem()).Elem().Interface()),
		}
	case reflect.Chan:
		panic("invalid type")
	case reflect.Func:
		panic("invalid type")
	case reflect.Interface:
		panic("invalid type")
	case reflect.Map:
		return &Type{
			IsMap:    true,
			ElemType: GetType(reflect.New(t.Elem()).Elem().Interface()),
		}
	case reflect.Pointer:
		return &Type{
			IsPointer: true,
			ElemType:  GetType(reflect.New(t.Elem()).Elem().Interface()),
		}
	case reflect.Slice:
		return &Type{
			IsArray:  true,
			ElemType: GetType(reflect.New(t.Elem()).Elem().Interface()),
		}
	case reflect.Struct:
		var fields []Field
		for i := 0; i < t.NumField(); i++ {
			f := t.Field(i)
			fields = append(fields, Field{
				ID:   f.Name,
				Name: f.Name,
				Type: GetType(reflect.New(f.Type).Elem().Interface()),
			})
		}
		return &Type{
			IsStruct: true,
			Fields:   fields,
		}
	case reflect.UnsafePointer:
		panic("invalid type")
	default:
		return &Type{
			IsScalar: true,
			Kind:     t.Kind().String(),
		}
	}
}

type Type struct {
	IsScalar  bool
	Kind      string
	IsPointer bool
	IsArray   bool
	IsMap     bool
	ElemType  *Type
	IsStruct  bool
	Fields    []Field
	Methods   map[string]*Function
}

func (t *Type) WriteGoAPI(w io.Writer) error {
	panic("not implemented")
}

func (t *Type) WriteNextJSPage(w io.Writer) error {
	el := &ReactElement{}
	if t.IsScalar {
		el.Type = t.Kind
	} else if t.IsPointer {
		el.Type = "Pointer"
	} else if t.IsArray {
		el.Type = "Array"
	} else if t.IsMap {
		el.Type = "Map"
	} else if t.IsStruct {
		el.Type = "Struct"
	}
	return el.WriteNextJSPage(w)
}
