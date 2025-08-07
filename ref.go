package util

import "fmt"

type Ref struct {
	From string `json:"from"`
	Name string `json:"name"`
}

func (r *Ref) GoImports() map[string]bool {
	if r.From == "" {
		return map[string]bool{}
	}
	return map[string]bool{
		r.From: true,
	}
}

func (r *Ref) GoString(imports map[string]string) string {
	if r.From == "" {
		return r.Name
	}
	return fmt.Sprintf("%s.%s", imports[r.From], r.Name)
}
