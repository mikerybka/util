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

func (t *Type) GoFile() []byte {
	panic("not implemented")
}
func (t *Type) DartFile() []byte {
	panic("not implemented")
}
func (t *Type) TypescriptFile() []byte {
	panic("not implemented")
}
func (t *Type) JavascriptFile() []byte {
	panic("not implemented")
}
