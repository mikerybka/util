package util

import (
	"fmt"
	"net/http"
)

type Link struct {
	Name string
	Href string
}

func (l *Link) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<a href=\"%s\">%s</a>", l.Href, l.Name)
}
