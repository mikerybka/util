package util

import (
	_ "embed"
)

type Object interface {
	ID() string
	JSON() string
	Type() string
	Ptr() any
	Dig(s string) Object
}

// //go:embed web_templates/object.html
// var objectTmpl string

// type Object[T any] struct {
// 	IsDir   bool
// 	Entries []string
// 	IsFile  bool
// 	Data    T
// }

// func (obj *Object[T]) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	if Accept(r, "text/html") {
// 		template.Must(template.New("object").Parse(objectTmpl)).Execute(w, obj)
// 		return
// 	}

// 	json.NewEncoder(w).Encode(obj)
// }
