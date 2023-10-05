package util

import (
	"encoding/json"
	"io"
)

func FprintJSON(w io.Writer, v any) (int, error) {
	b, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return 0, err
	}
	b = append(b, '\n')
	return w.Write(b)
}
