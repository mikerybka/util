package util

type Field struct {
	Name        Name   `json:"name"`
	Description string `json:"description"`
	Type        *Ref   `json:"type"`
}

func (f *Field) ID() string {
	return f.Name.ID()
}

func (f *Field) GoImports() map[string]bool {
	return f.Type.GoImports()
}
