package util

import (
	"net/http"

	"golang.org/x/net/html"
)

type Form struct {
	Name      string
	Fields    []Field
	ServePOST func(w http.ResponseWriter, r *http.Request)
}

func (f *Form) HTML() *html.Node {
	node := &html.Node{
		Type: html.ElementNode,
		Data: "form",
		Attr: []html.Attribute{
			{
				Key: "action",
				Val: ".",
			},
			{
				Key: "method",
				Val: "post",
			},
		},
	}
	for _, field := range f.Fields {
		label := &html.Node{
			Type: html.ElementNode,
			Data: "label",
			Attr: []html.Attribute{
				{
					Key: "for",
					Val: "name",
				},
			},
		}
		label.AppendChild(&html.Node{
			Type: html.TextNode,
			Data: field.Name + ":",
		})
		node.AppendChild(label)
		node.AppendChild(&html.Node{
			Type: html.ElementNode,
			Data: "input",
			Attr: []html.Attribute{
				{
					Key: "type",
					Val: "text",
				},
				{
					Key: "id",
					Val: field.Name,
				},
				{
					Key: "name",
					Val: field.Name,
				},
			},
		})
		node.AppendChild(&html.Node{
			Type: html.ElementNode,
			Data: "br",
		})
	}
	node.AppendChild(&html.Node{
		Type: html.ElementNode,
		Data: "input",
		Attr: []html.Attribute{
			{
				Key: "type",
				Val: "submit",
			},
			{
				Key: "value",
				Val: "Submit",
			},
		},
	})
	return node
}

func (f *Form) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		doc := &HTMLDocument{
			Head: &HTMLHead{
				Title: f.Name,
			},
			Body: f.HTML(),
		}
		doc.Write(w)
	} else if r.Method == http.MethodPost {
		f.ServePOST(w, r)
	}
}
