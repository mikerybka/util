package util

import (
	_ "embed"
	"os"
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

func (t *StructType) WriteGoFile(path, pkgName string) error {
	f := &GoFile{}
	return os.WriteFile(path, []byte(f.String()), os.ModePerm)
}
