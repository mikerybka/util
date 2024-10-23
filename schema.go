package util

type Schema struct {
	Fields []Field
}

func (s *Schema) Type() *Type {
	return &Type{
		IsStruct: true,
		Fields:   s.Fields,
	}
}
