package util

import "strings"

func OnlyLowerCase(s string) string {
	res := ""
	for _, r := range strings.ToLower(s) {
		if strings.ContainsRune("abcdefghijklmnopqrstuvwxyz", r) {
			res += string(r)
		}
	}
	return res
}
