package util

import (
	_ "embed"
)

type StructType struct {
	Name        Name                 `json:"name"`
	Description string               `json:"description"`
	Fields      []Field              `json:"fields"`
	Methods     map[string]*Function `json:"methods"`
	DefaultJSON string               `json:"default_json"`
}

//go:embed struct_type.ts.tmpl
var tsStructTypeTmpl string

func (t *StructType) WriteTypeScriptFile(path string) error {
	return RenderTemplateToFile(tsStructTypeTmpl, path, t)
}

//go:embed struct_type.go.tmpl
var goStructTypeTmpl string

func (t *StructType) WriteGoFile(path string) error {
	return RenderTemplateToFile(goStructTypeTmpl, path, t)
}
