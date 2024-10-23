package util

import (
	_ "embed"
)

type ScalarType struct {
	Name        Name                 `json:"name"`
	Description string               `json:"description"`
	ElemType    string               `json:"elemType"`
	Methods     map[string]*Function `json:"methods"`
	DefaultJSON string               `json:"defaultJSON"`
}

//go:embed scalar_type.ts.tmpl
var tsScalarTypeTmpl string

func (st *ScalarType) WriteTypeScriptFile(path string) error {
	return RenderTemplateToFile(tsScalarTypeTmpl, path, st)
}

//go:embed scalar_type.go.tmpl
var goScalarTypeTmpl string

func (st *ScalarType) WriteGoFile(path string) error {
	return RenderTemplateToFile(goScalarTypeTmpl, path, st)
}
