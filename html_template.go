package util

import "html/template"

func htmlTmpl(name, tmpl string) *template.Template {
	return template.Must(template.New(name).Parse(tmpl))

}
