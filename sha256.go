package util

import (
	"crypto/sha256"
	"encoding/hex"
)

func SHA256(b []byte) string {
	h := sha256.New()
	_, err := h.Write(b)
	if err != nil {
		panic(err)
	}
	hash := hex.EncodeToString(h.Sum(nil))
	return hash
}
