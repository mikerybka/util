package util

import (
	_ "embed"
	"html/template"
	"net/http"
)

//go:embed web/dist/main.js
var mainJS []byte

//go:embed web/dist/main.css
var mainCSS []byte

var htmlTemplate string = `<!DOCTYPE html>
<html>
	<head>
		<meta name="viewport" content="width=device-width, initial-scale=1.0" />
		<title>{{ .Title }}</title>
		<meta name="description" content="{{ .Desc }}"/>
		<meta name="author" content="{{ .Author }}" />
		<meta name="keywords" content="{{ .Keywords }}" />
		<link rel="stylesheet" href="/main.css" type="text/css" />
  	</head>
	<body>
		<main id="root">Loading...</main>
		<script src="/main.js"></script>
	</body>
</html>`

type WebFrontend struct {
	Favicon      []byte
	RootTitle    string
	MetaDesc     string
	MetaAuthor   string
	MetaKeywords []string
}

func (fe *WebFrontend) HTMLData() *HTMLTemplateData {
	keywords := ""
	for i, kw := range fe.MetaKeywords {
		if i > 0 {
			keywords += ","
		}
		keywords += kw
	}
	return &HTMLTemplateData{
		Title:    fe.RootTitle,
		Desc:     fe.MetaDesc,
		Author:   fe.MetaAuthor,
		Keywords: keywords,
	}
}

func (fe *WebFrontend) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/favicon.ico" {
		w.Write(fe.Favicon)
		return
	}
	if r.URL.Path == "/main.js" {
		w.Write(mainJS)
		return
	}
	if r.URL.Path == "/main.css" {
		w.Write(mainCSS)
		return
	}

	err := template.Must(template.New("main.html").Parse(htmlTemplate)).Execute(w, fe.HTMLData())
	if err != nil {
		panic(err)
	}
}
