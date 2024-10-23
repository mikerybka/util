package util

import (
	_ "embed"
)

type MapType struct {
	Name        Name                 `json:"name"`
	Description string               `json:"description"`
	ElemType    string               `json:"elem_type"`
	Methods     map[string]*Function `json:"methods"`
	DefaultJSON string               `json:"default_json"`
}

//go:embed map_type.ts.tmpl
var tsMapTypeTmpl string

func (t *MapType) WriteTypeScriptFile(path string) error {
	return RenderTemplateToFile(tsMapTypeTmpl, path, t)
}

//go:embed map_type.go.tmpl
var goMapTypeTmpl string

func (t *MapType) WriteGoFile(path string) error {
	return RenderTemplateToFile(goMapTypeTmpl, path, t)
}
