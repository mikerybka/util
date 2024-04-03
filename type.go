package util

import "io"

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
