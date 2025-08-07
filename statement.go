package util

// Statement represents a line of code in a function or method body.
// There are 4 types of statements: returns, assignments, ifs and loops.
// Ifs and loops have substatements.
type Statement struct {
	IsReturn bool        `json:"isReturn"`
	Return   *Expression `json:"return"`

	IsAssign bool        `json:"isAssign"`
	Name     string      `json:"name"`
	Value    *Expression `json:"value"`

	IsIf      bool         `json:"isIf"`
	Condition *Expression  `json:"condition"`
	Body      []*Statement `json:"body"`

	// TODO: loops
}

func (st *Statement) GoString(imports map[string]string, indent int) string {
	s := ""
	for i := 0; i < indent; i++ {
		s += "\t"
	}
	if st.IsReturn {
		s += "return "
		s += st.Return.Value
		s += "\n"
	}
	if st.IsAssign {
		s += st.Name
		s += " = "
		s += st.Value.GoString(imports)
		s += "\n"
	}
	if st.IsIf {
		s += "if "
		s += st.Condition.GoString(imports)
		s += " {\n"
		for _, st := range st.Body {
			s += st.GoString(imports, indent+1)
		}
		s += "}\n"
	}
	return s
}

func (s *Statement) GoImports() map[string]bool {
	imports := map[string]bool{}
	if s.IsReturn {
		for k, v := range s.Return.GoImports() {
			if v {
				imports[k] = true
			}
		}
	}
	if s.IsAssign {
		for k, v := range s.Value.GoImports() {
			if v {
				imports[k] = true
			}
		}
	}
	if s.IsIf {
		for _, st := range s.Body {
			for k, v := range st.GoImports() {
				if v {
					imports[k] = true
				}
			}
		}
	}
	return imports
}
