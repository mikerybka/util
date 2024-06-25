package util

import (
	"fmt"
	"io"
	"strings"
)

type Type struct {
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

func (t *Type) GoString(indent int) string {
	if t.IsScalar {
		return t.Kind
	}
	if t.IsPointer {
		return "*" + t.ElemType
	}
	if t.IsArray {
		return "[]" + t.ElemType
	}
	if t.IsMap {
		return "map[string]" + t.ElemType
	}
	if t.IsStruct {
		s := strings.Builder{}
		s.WriteString("struct {\n")
		for _, f := range t.Fields {
			for i := 0; i < indent+1; i++ {
				s.WriteString("\t")
			}
			s.WriteString(f.Name)
			s.WriteString(" ")
			s.WriteString(f.Type.GoString(indent + 1))
			s.WriteString("\n")
		}
		s.WriteString("}")
		return s.String()
	}
	panic("unknown type")
}

func (t *Type) WriteGoFile(w io.Writer, pkg, name string) (int, error) {
	return fmt.Fprintf(w, "package %s\n\ntype %s %s\n", pkg, name, t.GoString(0))
}

func (t *Type) GoFile(pkg, typ string) string {
	return fmt.Sprintf("package %s\n\ntype %s %s\n", pkg, typ, t.GoString(0))
}
