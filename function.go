package util

import (
	_ "embed"
)

type Function struct {
	Name        Name        `json:"name"`
	Description string      `json:"description"`
	Inputs      []Field     `json:"inputs"`
	Outputs     []Field     `json:"outputs"`
	Body        []Statement `json:"body"`
}

//go:embed function.go.tmpl
var goFuncTmpl string

func (f *Function) WriteGoFile(path string) error {
	return RenderTemplateToFile(goFuncTmpl, path, f)
}

//go:embed function.ts.tmpl
var tsFuncTmpl string

func (f *Function) WriteTypeScriptFile(path string) error {
	return RenderTemplateToFile(tsFuncTmpl, path, f)
}
