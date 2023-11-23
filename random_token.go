package util

import (
	"crypto/rand"
	"encoding/hex"
)

func RandomToken(bytes int) string {
	b := make([]byte, bytes)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}
	return hex.EncodeToString(b)
}
