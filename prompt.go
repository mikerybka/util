package util

import (
	_ "embed"
	"html/template"
	"net/http"
)

type Prompt struct {
	Name    Name
	Desc    string
	Options []PromptOption
}

type PromptOption struct {
	Name Name
	URL  string
}

func (p *Prompt) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := template.Must(template.New("prompt").Parse(promptHTML)).Execute(w, p)
	if err != nil {
		panic(err)
	}
}

//go:embed prompt.html
var promptHTML string
