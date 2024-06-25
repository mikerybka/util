package util

import (
	"io"
	"net/http"
)

type HTMLDocument struct {
	Head *HTMLHead
	Body *XML
}

func (d *HTMLDocument) XML() *XML {
	return &XML{
		El: "html",
		Children: []*XML{
			d.Head.XML(),
			d.Body,
		},
	}
}

func (d *HTMLDocument) Write(w io.Writer) (int64, error) {
	var written int64
	n32, err := w.Write([]byte("<!DOCTYPE html>"))
	written += int64(n32)
	if err != nil {
		return written, err
	}
	n, err := d.XML().Write(w)
	written += n
	return written, err
}

func (d *HTMLDocument) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	d.Write(w)
}
