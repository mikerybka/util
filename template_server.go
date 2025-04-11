package util

import (
	"html/template"
	"net/http"
	"path/filepath"
)

func NewTemplateServer[DataType any](tmpl *template.Template, workdir, userID string) *TemplateServer[DataType] {
	return &TemplateServer[DataType]{
		Tmpl:    tmpl,
		Workdir: workdir,
		UserID:  userID,
	}
}

type TemplateServer[DataType any] struct {
	Tmpl    *template.Template
	Workdir string
	UserID  string
}

func (s *TemplateServer[T]) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	d := new(T)
	path := filepath.Join(s.Workdir, r.URL.Path)
	err := ReadJSONFile(path, d)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// TODO: add func for PathValue
	err = s.Tmpl.Execute(w, struct {
		UserID string
		Data   any
	}{
		UserID: s.UserID,
		Data:   d,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
