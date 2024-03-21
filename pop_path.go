package util

import "strings"

func PopPath(path string) string {
	p := ParsePath(path)
	return "/" + strings.Join(p[1:], "/")
}
