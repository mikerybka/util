package util

import (
	"encoding/json"
	"io"
)

func ReadJSON[T any](r io.Reader) T {
	var v T
	if err := json.NewDecoder(r).Decode(&v); err != nil {
		panic(err)
	}
	return v
}
