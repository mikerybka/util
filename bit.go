package util

import "fmt"

// Bit returns the bit value of the bit at index i of b.
// It panics if i is out of range (1-8).
func Bit(b byte, i int) bool {
	bits := fmt.Sprintf("%08b", b)
	switch bits[7-i] {
	case '0':
		return false
	case '1':
		return true
	default:
		panic("unreachable")
	}
}
