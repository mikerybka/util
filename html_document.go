package util

import (
	"io"

	"golang.org/x/net/html"
)

type HTMLDocument struct {
	Head *HTMLHead
	Body *html.Node
}

func (d *HTMLDocument) Write(w io.Writer) error {
	w.Write([]byte("<!DOCTYPE html>"))
	w.Write([]byte("<html>"))
	d.Head.Write(w)
	w.Write([]byte("<body>"))
	err := html.Render(w, d.Body)
	w.Write([]byte("</body>"))
	w.Write([]byte("</html>"))
	return err
}
