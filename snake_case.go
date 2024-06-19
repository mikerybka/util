package util

import "strings"

func SnakeCase(s string) string {
	return strings.ReplaceAll(strings.ToLower(s), " ", "_")
}
