package util

import (
	"golang.org/x/net/html"
)

type HTMLDocument struct {
	Head *HTMLHead
	Body *html.Node
}
