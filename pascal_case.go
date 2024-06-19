package util

import (
	"strings"
)

func PascalCase(s string) string {
	return strings.ReplaceAll(TitleCase(s), " ", "")
}
