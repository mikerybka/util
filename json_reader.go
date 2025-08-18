package util

import "io"

func NewJSONReader(types map[string]*Type) *JSONReader {
	return &JSONReader{
		Types: types,
	}
}

type JSONReader struct {
	Types map[string]*Type
}

func (r *JSONReader) Read(w io.Writer, t, path string) error
