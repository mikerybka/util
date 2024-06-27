package util

import (
	"bytes"
	"io"

	"golang.org/x/net/html"
)

func NewXML(n *html.Node) *XML {
	return &XML{
		El: n.Data,
	}
}

type XML struct {
	El       string
	Attrs    []Pair[string, string]
	Children []*XML
}

func (xml *XML) String() string {
	buf := &bytes.Buffer{}
	_, err := xml.Write(buf)
	if err != nil {
		panic(err)
	}
	return buf.String()
}

func (xml *XML) Attr(key string) string {
	for _, attr := range xml.Attrs {
		if attr.K == key {
			return attr.V
		}
	}
	return ""
}

func (xml *XML) Write(w io.Writer) (n int64, err error) {
	buf := &bytes.Buffer{}

	// Handle strings
	if xml.El == "string" {
		buf.WriteString(xml.Attr("value"))
		return buf.WriteTo(w)
	}

	// Open
	buf.WriteString("<")
	buf.WriteString(xml.El)
	for _, attr := range xml.Attrs {
		buf.WriteString(" ")
		buf.WriteString(attr.K)
		buf.WriteString("=\"")
		buf.WriteString(attr.V)
		buf.WriteString("\"")
	}
	buf.WriteString(">")

	// Children
	for _, ch := range xml.Children {
		buf.WriteString(ch.String())
	}

	// Close
	buf.WriteString("</")
	buf.WriteString(xml.El)
	buf.WriteString(">")

	// Write out
	return buf.WriteTo(w)
}
