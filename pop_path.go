package util

import "strings"

func PopPath(path string) (first string, rest string, isRoot bool) {
	p := ParsePath(path)
	if len(p) == 0 {
		return "", "/", true
	}
	return p[0], "/" + strings.Join(p[1:], "/"), false
}
