package util

import "golang.org/x/net/html"

type Action struct {
	Name string
	URL  string
}

func (a *Action) HTML() *html.Node {
	link := &html.Node{
		Type: html.ElementNode,
		Data: "a",
		Attr: []html.Attribute{
			{
				Key: "href",
				Val: a.URL,
			},
		},
	}
	txt := &html.Node{
		Type: html.TextNode,
		Data: a.Name,
	}
	link.AppendChild(txt)
	return link
}
