package util

import (
	"net/http"
	"text/template"
)

type LinkList struct {
	Links []Link
}

func (s *LinkList) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	template.Must(template.New("t").Parse(`<pre>{{ range .Links }}
<a href="{{ .Href }}">{{ .Name }}</a>{{ end }}</pre>`))
}
