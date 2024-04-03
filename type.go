package util

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
