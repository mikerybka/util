package util

import "strings"

func JoinPath(path []string) string {
	return "/" + strings.Join(path, "/")
}
