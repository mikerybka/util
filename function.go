package util

import (
	_ "embed"
)

type Function struct {
	Description string      `json:"description"`
	Inputs      []Field     `json:"inputs"`
	Outputs     []Field     `json:"outputs"`
	Body        []Statement `json:"body"`
}

func (f *Function) GoString(imports map[string]string) string {
	s := "("
	for i, input := range f.Inputs {
		if i > 0 {
			s += ", "
		}
		s += input.Name.SnakeCase()
		s += " "
		s += input.Type.GoString(imports)
	}
	s += ")"

	if len(f.Outputs) > 0 {
		s += " "
	}
	if len(f.Outputs) > 1 {
		s += "("
	}
	for i, o := range f.Outputs {
		if i > 0 {
			s += ", "
		}
		s += o.Name.SnakeCase()
		s += " "
		s += o.Type.GoString(imports)
	}
	if len(f.Outputs) > 1 {
		s += ")"
	}

	s += " {\n"
	for _, st := range f.Body {
		s += st.GoString(imports, 1)
	}
	s += "}\n"
	return s
}

func (f *Function) GoImports() map[string]bool {
	imports := map[string]bool{}
	for _, in := range f.Inputs {
		for k, v := range in.GoImports() {
			if v {
				imports[k] = true
			}
		}
	}
	for _, out := range f.Outputs {
		for k, v := range out.GoImports() {
			if v {
				imports[k] = true
			}
		}
	}
	for _, st := range f.Body {
		for k, v := range st.GoImports() {
			if v {
				imports[k] = true
			}
		}
	}
	return imports
}
