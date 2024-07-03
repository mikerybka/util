package util

import (
	"encoding/json"
	"io"
)

func ReadJSON[T any](r io.Reader) T {
	var v T
	json.NewDecoder(r).Decode(v)
	return v
}
