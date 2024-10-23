package util

import (
	_ "embed"
)

type ReactComponent struct {
	Name  Name
	Props []*Field
	State []*Field
	Body  []*Statement
}

//go:embed react_component.tsx.tmpl
var reactComponentTmpl string

func (rc *ReactComponent) Write(path string) error {
	return RenderTemplateToFile(reactComponentTmpl, path, rc)
}
