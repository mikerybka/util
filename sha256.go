package util

import (
	"crypto/sha256"
)

func SHA256(b []byte) []byte {
	h := sha256.New()
	_, err := h.Write(b)
	if err != nil {
		panic(err)
	}
	return h.Sum(nil)
}
