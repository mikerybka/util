package util

import "strings"

func ParsePath(s string) []string {
	path := []string{}
	for _, p := range strings.Split(s, "/") {
		if p != "" {
			path = append(path, p)
		}
	}
	return path
}
