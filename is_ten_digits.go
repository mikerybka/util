package util

import "unicode"

func IsTenDigits(phone string) bool {
	if len(phone) != 10 {
		return false
	}
	for _, ch := range phone {
		if !unicode.IsDigit(ch) {
			return false
		}
	}
	return true
}
