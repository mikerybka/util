package util

import "strings"

func PopPath(path string) (string, string, bool) {
	p := ParsePath(path)
	if len(p) == 0 {
		return "", "", false
	}
	return p[0], "/" + strings.Join(p[1:], "/"), true
}
