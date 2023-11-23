package util

import (
	"fmt"
	"strings"
	"unicode"
)

func NormalizeGoName(s string) string {
	runes := []rune{}
	for _, r := range s {
		runes = append(runes, r)
	}
	words := []string{}
	isAcro := false
	for i := range runes {
		ch := runes[i]
		if unicode.IsUpper(ch) {
			if isAcro {
				if i == len(runes)-1 || unicode.IsUpper(runes[i+1]) {
					words[len(words)-1] += string(unicode.ToLower(ch))
				} else {
					words = append(words, string(unicode.ToLower(ch)))
				}
			} else {
				words = append(words, string(unicode.ToLower(ch)))
				isAcro = true
			}
		} else {
			words[len(words)-1] += string(unicode.ToLower(ch))
			isAcro = false
		}
	}
	for i := range words {
		stripped := StripNonAlphaNumeric(words[i])
		if stripped == "" {
			panic(fmt.Errorf("bad word: %s", words[i]))
		}
		words[i] = stripped
	}
	return strings.Join(words, "-")
}
