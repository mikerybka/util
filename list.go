package util

import (
	_ "embed"
	"net/http"
)

type List struct {
	Size  int
	Items []ListItem
}

//go:embed list.html
var listHTML string

func (l *List) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	htmlTmpl("list", listHTML).Execute(w, l)
}
