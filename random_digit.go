package util

import "crypto/rand"

func RandomDigit() int {
	// r := RandomBits(4)
	randomByte := make([]byte, 1)
	_, err := rand.Read(randomByte)
	if err != nil {
		panic(err)
	}
	i := int(randomByte[0]) % 16
	if i < 10 {
		return i
	}
	return RandomDigit()
}
