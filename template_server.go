package util

import (
	"html/template"
	"net/http"
)

func NewTemplateServer[DataType any](tmpl *template.Template, dataPath string) *TemplateServer[DataType] {
	return &TemplateServer[DataType]{
		Tmpl:     tmpl,
		DataPath: dataPath,
	}
}

type TemplateServer[DataType any] struct {
	Tmpl     *template.Template
	DataPath string
}

func (s *TemplateServer[T]) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	d := new(T)
	err := ReadJSONFile(s.DataPath, d)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	s.Tmpl.Execute(w, d)
}
