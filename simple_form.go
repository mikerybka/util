package util

import (
	_ "embed"
	"net/http"
	"text/template"
)

type SimpleForm struct {
	TitleText  string
	Fields     []Field
	SubmitText string
	Error      error
	HandlePOST http.HandlerFunc
}

func (form *SimpleForm) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		simpleFormTmpl.Execute(w, form)
		return
	}
	if r.Method == "POST" {
		form.HandlePOST(w, r)
		return
	}
	http.NotFound(w, r)
}

//go:embed simple_form.html
var simpleFormHTML string
var simpleFormTmpl = template.Must(template.New("simple-form").Parse(simpleFormHTML))
