package util

import (
	"math/rand/v2"
)

func RandomString(length int, charset string) string {
	s := make([]byte, length)
	for i := 0; i < length; i++ {
		s[i] = charset[rand.IntN(len(charset))]
	}
	return string(s)
}
