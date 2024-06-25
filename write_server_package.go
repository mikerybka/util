package util

import (
	"bytes"
	"fmt"
	"path/filepath"
	"text/template"
)

func WriteServerPackage(schemas map[string]Schema, dir string) error {
	// Write pkg/types
	for id, s := range schemas {
		filename := filepath.Join(dir, "pkg/types", SnakeCase(id)+".go")
		contents := s.Type().GoFile(OnlyLowerCase(id), PascalCase(id))
		err := WriteFile(filename, []byte(contents))
		if err != nil {
			return fmt.Errorf("writing %s: %s", filename, err)
		}
	}

	// Write main.go
	t := template.Must(template.New("main.go").Parse(MainGoServerTemplate))
	buf := bytes.NewBuffer(nil)
	err := t.Execute(buf, schemas)
	if err != nil {
		panic(err)
	}
	contents := buf.Bytes()
	filename := filepath.Join(dir, "main.go")
	err = WriteFile(filename, contents)
	if err != nil {
		return fmt.Errorf("writing main.go: %s", err)
	}

	return nil
}
