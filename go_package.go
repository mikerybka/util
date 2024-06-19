package util

import (
	"os"
	"path/filepath"
)

type GoPackage struct {
	Types map[string]Type
}

func (p *GoPackage) Write(dir string) error {
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return err
	}

	for name, t := range p.Types {
		filename := SnakeCase(name) + ".go"
		typeName := PascalCase(name)
		path := filepath.Join(dir, filename)
		f, err := os.Create(path)
		if err != nil {
			return err
		}
		pkgName := filepath.Base(dir)
		t.WriteGoFile(f, pkgName, typeName)
	}

	return nil
}
