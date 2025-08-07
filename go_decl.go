package util

import "fmt"

type GoDecl struct {
	Name       Name      `json:"name"`
	IsType     bool      `json:"isType"`
	Type       *Type     `json:"type"`
	IsFunction bool      `json:"isFunction"`
	Function   *Function `json:"function"`
	IsVar      bool      `json:"isVar"`
	Var        *Value    `json:"var"`
	IsConst    bool      `json:"isConst"`
	Const      *Value    `json:"const"`
}

func (decl *GoDecl) GoImports() map[string]bool {
	if decl.IsConst {
		return decl.Const.GoImports()
	}
	if decl.IsVar {
		return decl.Var.GoImports()
	}
	if decl.IsFunction {
		return decl.Function.GoImports()
	}
	if decl.IsType {
		return decl.Type.GoImports()
	}
	panic("bad decl")
}

func (decl *GoDecl) String(imports map[string]string) string {
	if decl.IsConst {
		return fmt.Sprintf("const %s = %s\n", decl.Name, decl.Const.GoString(imports))
	}
	if decl.IsVar {
		return fmt.Sprintf("var %s = %s\n", decl.Name, decl.Var.GoString(imports))
	}
	if decl.IsFunction {
		return fmt.Sprintf("func %s%s", decl.Name, decl.Function.GoString(imports))
	}
	if decl.IsType {
		s := fmt.Sprintf("type %s %s", decl.Name, decl.Type.GoString(imports))
		for _, name := range SortedKeys(decl.Type.Methods) {
			m := decl.Type.Methods[name]
			s += fmt.Sprintf("\nfunc (%s *%s) %s%s", decl.Type.LocalInstanceName, decl.Name.GoExported(), name, m.GoString(imports))
		}
		return s
	}
	panic("bad decl")
}
