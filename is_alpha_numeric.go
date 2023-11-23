package util

import "unicode"

func IsAlphaNumeric(ch rune) bool {
	return (unicode.IsDigit(ch) || unicode.IsLetter(ch)) && unicode.IsPrint(ch)
}
