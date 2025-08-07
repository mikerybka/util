package util

type Expression struct {
	IsLiteral bool   `json:"isLiteral"`
	Value     string `json:"value"`

	IsCall bool          `json:"isCall"`
	Fn     *Ref          `json:"fn"`
	Args   []*Expression `json:"args"`

	IsRef bool `json:"isRef"`
	Ref   *Ref `json:"ref"`
}

func (ex *Expression) GoString(imports map[string]string) string {
	if ex.IsLiteral {
		return ex.Value
	}
	if ex.IsCall {
		s := ex.Fn.GoString(imports)
		s += "("
		for i, arg := range ex.Args {
			if i > 0 {
				s += ", "
			}
			s += arg.GoString(imports)
		}
		s += ")"
		return s
	}
	if ex.IsRef {
		return ex.Ref.GoString(imports)
	}
	panic("bad expr")
}

func (ex *Expression) GoImports() map[string]bool {
	imports := map[string]bool{}
	for _, arg := range ex.Args {
		for k, v := range arg.GoImports() {
			if v {
				imports[k] = true
			}
		}
	}
	return imports
}
