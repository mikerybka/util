package util

import "strings"

func KebabCase(s string) string {
	validChars := "abcdefghijklmnopqrstuvwxyz"
	newString := ""
	for _, ch := range strings.ToLower(s) {
		if ch == ' ' {
			newString += "-"
		} else if strings.Contains(validChars, string(ch)) {
			newString += string(ch)
		}
	}
	return newString
}
