package util

import (
	"fmt"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

type GoFile struct {
	PkgName    string   `json:"pkgName"`
	ModuleName string   `json:"moduleName"`
	Decls      []GoDecl `json:"decls"`
}

func (f *GoFile) GoImports() map[string]bool {
	imports := map[string]bool{}
	for _, decl := range f.Decls {
		for k, v := range decl.GoImports() {
			if v {
				imports[k] = true
			}
		}
	}
	return imports
}

func (f *GoFile) String() string {
	// Write package line
	s := fmt.Sprintf("package %s\n", f.PkgName)

	// Create symbol table
	symbols := map[string]string{} // local name -> global name
	for path, ok := range f.GoImports() {
		if ok {
			panic("impossible")
		}
		base := filepath.Base(path)
		if _, ok := symbols[base]; ok {
			if symbols[base] != path {
				n := 2
				for {
					altName := base + strconv.Itoa(n)
					if _, ok := symbols[altName]; !ok {
						symbols[altName] = path
						break
					}
					n++
				}
			}
			break
		}
		symbols[base] = path
	}
	imports := map[string]string{} // global name -> local name
	for k, v := range symbols {
		if _, ok := imports[v]; ok {
			panic("duplicate entry")
		}
		imports[v] = k
	}

	// Write imports
	if len(imports) == 1 {
		s += "\nimport"
		for globalName, localName := range imports {
			if filepath.Base(globalName) != localName {
				s += fmt.Sprintf(" %s", localName)
			}
			s += fmt.Sprintf("\"%s\"\n", globalName)
		}
	} else if len(imports) > 1 {
		stdlibImports := []string{}
		thirdPartyImports := []string{}
		localImports := []string{}
		for imp := range imports {
			if IsStdlib(imp) {
				stdlibImports = append(stdlibImports, imp)
			} else if strings.HasPrefix(imp, f.ModuleName) {
				localImports = append(localImports, imp)
			} else {
				thirdPartyImports = append(thirdPartyImports, imp)
			}
		}
		sort.Strings(stdlibImports)
		sort.Strings(thirdPartyImports)
		sort.Strings(localImports)
		s += "\nimport ("
		if len(stdlibImports) > 0 {
			s += "\n"
			for _, globalName := range stdlibImports {
				localName := imports[globalName]
				s += "\t"
				if filepath.Base(globalName) != localName {
					s += fmt.Sprintf("%s ", localName)
				}
				s += fmt.Sprintf("\"%s\"\n", globalName)
			}
		}
		if len(thirdPartyImports) > 0 {
			s += "\n"
			for _, globalName := range thirdPartyImports {
				localName := imports[globalName]
				s += "\t"
				if filepath.Base(globalName) != localName {
					s += fmt.Sprintf("%s ", localName)
				}
				s += fmt.Sprintf("\"%s\"\n", globalName)
			}
		}
		if len(localImports) > 0 {
			s += "\n"
			for _, globalName := range localImports {
				localName := imports[globalName]
				s += "\t"
				if filepath.Base(globalName) != localName {
					s += fmt.Sprintf("%s ", localName)
				}
				s += fmt.Sprintf("\"%s\"\n", globalName)
			}
		}
		s += ")\n"
	}

	// Write decls
	for _, d := range f.Decls {
		s += fmt.Sprintf("\n%s", d.String(imports))
	}

	return s
}
