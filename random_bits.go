package util

import "crypto/rand"

func RandomBits(n int) []bool {
	randomBytes := make([]byte, (n/8)+1)
	_, err := rand.Read(randomBytes)
	if err != nil {
		panic(err)
	}
	randomBits := make([]bool, n)
	for i := range randomBits {
		randomBits[i] = Bit(randomBytes[i/8], i%8)
	}
	return randomBits
}
