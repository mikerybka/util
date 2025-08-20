package util

type Type struct {
	Description       string               `json:"description"`
	IsScalar          bool                 `json:"isScalar"`
	Scalar            *Ref                 `json:"scalar"`
	IsArray           bool                 `json:"isArray"`
	IsMap             bool                 `json:"isMap"`
	ElemType          *Type                `json:"elemType"`
	IsStruct          bool                 `json:"isStruct"`
	Fields            []Field              `json:"fields"`
	Methods           map[string]*Function `json:"methods"`
	DefaultJSON       string               `json:"defaultJSON"`
	LocalInstanceName string               `json:"localInstanceName"`
}

func (t *Type) longestFieldName() int {
	n := 0
	for _, f := range t.Fields {
		name := f.Name.GoExported()
		l := len(name)
		if l > n {
			n = l
		}
	}
	return n
}

func (t *Type) GoString(imports map[string]string) string {
	if t.IsScalar {
		return t.Scalar.GoString(imports)
	} else if t.IsArray {
		return "[]" + t.ElemType.GoString(imports)
	} else if t.IsMap {
		return "map[string]" + t.ElemType.GoString(imports)
	} else if t.IsStruct {
		nameWidth := t.longestFieldName()
		s := "struct {\n"
		for _, f := range t.Fields {
			s += "\t"
			name := f.Name.GoExported()
			s += name
			for i := 0; i < 1+nameWidth-len(name); i++ {
				s += " "
			}
			s += f.Type.GoString(imports)
			s += "\n"
		}
		s += "}\n"
		return s
	} else {
		panic("bad type")
	}
}

func (t *Type) GoImports() map[string]bool {
	imports := t.ElemType.GoImports()
	for _, f := range t.Fields {
		for k, v := range f.GoImports() {
			if v {
				imports[k] = true
			}
		}
	}
	for _, m := range t.Methods {
		for k, v := range m.GoImports() {
			if v {
				imports[k] = true
			}
		}
	}
	return imports
}
