package util

import (
	_ "embed"
	"html/template"
	"net/http"
)

type Form struct {
	Name   Name
	Desc   string
	Fields []Field
	Handle http.HandlerFunc
}

func (f *Form) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		f.get(w, r)
	} else if r.Method == http.MethodPost {
		f.post(w, r)
	}
}

func (f *Form) get(w http.ResponseWriter, r *http.Request) {
	err := template.Must(template.New("form").Parse(formHTML)).Execute(w, f)
	if err != nil {
		panic(err)
	}
}

func (f *Form) post(w http.ResponseWriter, r *http.Request) {
	f.Handle(w, r)
}

//go:embed form.html
var formHTML string
