package util

import (
	"text/template"
)

func RenderTemplateToFile(tmpl string, path string, in any) error {
	f, err := CreateFile(path)
	if err != nil {
		return err
	}
	defer f.Close()
	return template.Must(template.New("t").Parse(tmpl)).Execute(f, in)
}
