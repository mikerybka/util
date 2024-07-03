package util

import (
	"golang.org/x/net/html"
)

type Error struct {
	Message string
	Actions []Action
}

func (e *Error) HTML() *html.Node {
	doc := &html.Node{
		Type: html.DoctypeNode,
		Data: "html",
	}
	htmlNode := &html.Node{
		Type: html.ElementNode,
		Data: "html",
	}
	doc.AppendChild(htmlNode)

	head := &html.Node{
		Type: html.ElementNode,
		Data: "head",
	}
	htmlNode.AppendChild(head)

	body := &html.Node{
		Type: html.ElementNode,
		Data: "body",
	}
	htmlNode.AppendChild(body)

	p := &html.Node{
		Type: html.ElementNode,
		Data: "p",
	}
	body.AppendChild(p)

	msg := &html.Node{
		Type: html.TextNode,
		Data: e.Message,
	}
	p.AppendChild(msg)

	actions := &html.Node{
		Type: html.ElementNode,
		Data: "div",
	}
	body.AppendChild(actions)

	for _, a := range e.Actions {
		actions.AppendChild(a.HTML())
	}

	return doc
}
