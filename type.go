package util

import (
	"io"
)

type Type struct {
	Name       string
	PluralName string

	IsScalar    bool
	Kind        string
	IsPointer   bool
	IsArray     bool
	IsMap       bool
	ElemType    string
	IsStruct    bool
	Fields      []Field
	Methods     map[string]*Function
	DefaultJSON string
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
