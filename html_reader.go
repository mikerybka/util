package util

import "io"

func NewHTMLReader(types map[string]*Type) *HTMLReader {
	return &HTMLReader{
		Types: types,
	}
}

type HTMLReader struct {
	Types map[string]*Type
}

func (r *HTMLReader) Read(w io.Writer, t, path string) error
