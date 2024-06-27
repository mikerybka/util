package main

import (
	"fmt"
	"strings"

	"golang.org/x/net/html"
)

// Traverse and print the HTML node tree
func traverse(n *html.Node) {
	if n.Type == html.ElementNode {
		fmt.Printf("<%s>\n", n.Data)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		traverse(c)
	}

	if n.Type == html.ElementNode {
		fmt.Printf("</%s>\n", n.Data)
	}
}

func main() {
	// Example HTML string
	htmlData := `<!DOCTYPE html>
	<html>
    <head>
        <title>Example</title>
    </head>
    <body>
        <h1>Hello, World!</h1>
        <p>This is a simple HTML example.</p>
    </body>
    </html>`

	// Parse the HTML string
	doc, err := html.Parse(strings.NewReader(htmlData))
	if err != nil {
		fmt.Printf("Error parsing HTML: %v\n", err)
		return
	}

	// Traverse and print the HTML node tree
	traverse(doc)
}
