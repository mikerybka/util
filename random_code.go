package util

import "strconv"

func RandomCode(digits int) string {
	s := ""
	for i := 0; i < digits; i++ {
		s += strconv.Itoa(RandomDigit())
	}
	return s
}
