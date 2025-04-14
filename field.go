package util

type Field struct {
	Name        Name   `json:"name"`
	Description string `json:"description"`
	Type        string `json:"type"`
}

func (f *Field) ID() string {
	return f.Name.ID()
}
