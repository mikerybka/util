package util

import (
	"fmt"
	"path/filepath"
)

// ImportMap maps symbols to Go import strings
type ImportMap map[string]string

func (im ImportMap) Go() string {
	if len(im) == 0 {
		return ""
	}

	if len(im) == 1 {
		for name, path := range im {
			if name == filepath.Base(path) {
				return fmt.Sprintf(`import "%s"`, path)
			} else {
				return fmt.Sprintf(`import %s "%s"`, name, path)
			}
		}
	}

	// len(im) > 1
	out := "import (\n"
	for name, path := range im {
		if name == filepath.Base(path) {
			out += fmt.Sprintf("\t\"%s\"\n", path)
		} else {
			out += fmt.Sprintf("\t%s \"%s\"\n", path)
		}
	}
	out += ")\n"
	return out
}

func (im ImportMap) TypeScript() string {
	out := ""
	for key, path := range im {
		out += fmt.Sprintf("import %s from %s;\n", key, path)
	}
	return out
}
