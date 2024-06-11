package util

import (
	"fmt"
	"net/http"
)

type MapList[T http.Handler] Map[T]

func (m MapList[T]) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<div class=\"list\">")
	for k, v := range m {
		l := &Link{
			Name: GetName(v),
			Href: k,
		}
		l.ServeHTTP(w, r)
	}
	fmt.Fprintf(w, "</div>")
}
