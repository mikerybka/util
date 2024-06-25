package util

import "io"

type Schema struct {
	Fields []Field
}

func (s *Schema) Type() *Type {
	return &Type{
		IsStruct: true,
		Fields:   s.Fields,
	}
}

func (s *Schema) WriteGoTypeFile(w io.Writer, pkg, name string) (int, error) {
	return s.Type().WriteGoFile(w, pkg, name)
}
