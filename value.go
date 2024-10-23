package util

import (
	_ "embed"
)

type Value struct {
	Name        Name   `json:"name"`
	Description string `json:"description"`
	Type        string `json:"type"`
	Data        any    `json:"data"`
}

//go:embed const.ts.tmpl
var tsConstTmpl string

func (v *Value) WriteTypeScriptConstantFile(path string) error {
	return RenderTemplateToFile(tsConstTmpl, path, v)
}

//go:embed react_hook.ts.tmpl
var reactHookTmpl string

func (v *Value) WriteReactHookFile(path string) error {
	return RenderTemplateToFile(reactHookTmpl, path, v)
}

//go:embed const.go.tmpl
var goConstTmpl string

func (v *Value) WriteGoConstFile(path string) error {
	return RenderTemplateToFile(goConstTmpl, path, v)
}

//go:embed var.go.tmpl
var goVarTmpl string

func (v *Value) WriteGoVarFile(path string) error {
	return RenderTemplateToFile(goVarTmpl, path, v)
}
