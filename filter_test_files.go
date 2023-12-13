package util

import (
	"io/fs"
	"strings"
)

func FilterTestFiles(fi fs.FileInfo) bool {
	return !strings.HasSuffix(fi.Name(), "_test.go")
}
