package util

import "strings"

func FilterString(input, charset string) string {
	allowed := make(map[rune]bool)
	for _, ch := range charset {
		allowed[ch] = true
	}

	var builder strings.Builder
	for _, ch := range input {
		if allowed[ch] {
			builder.WriteRune(ch)
		}
	}

	return builder.String()
}
