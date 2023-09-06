package util

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"strings"
)

func filterTests(f os.FileInfo) bool {
	return !strings.HasSuffix(f.Name(), "_test.go")
}

func ReadGoPackageFromDisk(path string) (*ast.Package, error) {
	fset := token.NewFileSet()
	pkgs, err := parser.ParseDir(fset, path, filterTests, parser.ParseComments)
	if err != nil {
		return nil, err
	}
	if len(pkgs) != 1 {
		return nil, fmt.Errorf("expected 1 package, got %d", len(pkgs))
	}
	for _, pkg := range pkgs {
		return pkg, nil
	}
	panic("unreachable")
}
