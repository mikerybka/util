package util

import (
	_ "embed"
)

type ArrayType struct {
	Name        Name                 `json:"name"`
	Description string               `json:"description"`
	ElemType    string               `json:"elem_type"`
	Methods     map[string]*Function `json:"methods"`
	DefaultJSON string               `json:"default_json"`
}

//go:embed array_type.ts.tmpl
var tsArrayTypeTmpl string

func (t *ArrayType) WriteTypeScriptFile(path string) error {
	return RenderTemplateToFile(tsArrayTypeTmpl, path, t)
}

//go:embed array_type.go.tmpl
var goArrayTypeTmpl string

func (t *ArrayType) WriteGoFile(path string) error {
	return RenderTemplateToFile(goArrayTypeTmpl, path, t)
}
