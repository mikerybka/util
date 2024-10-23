package util

import (
	"fmt"
	"os/exec"
	"path/filepath"
)

type NextApp struct {
	Favicon     []byte
	StaticFiles map[string][]byte
	Types       map[string]Type
	Constants   map[string]Value
	Functions   map[string]Function
	Hooks       map[string]Function
	Components  map[string]ReactComponent
	Pages       map[string]NextPage
	Layouts     map[string]ReactComponent
}

func (na *NextApp) Write(dir string) error {
	// Generate with `npx create-next-app@latest`
	cmd := exec.Command("npx", "create-next-app@latest", "--typescript", "--eslint", "--tailwind", "--src-dir", "--app", filepath.Base(dir))
	cmd.Dir = filepath.Dir(dir)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("%s%s", out, err)
	}

	// Write favicon.ico
	path := filepath.Join(dir, "web/src/app/favicon.ico")
	err = WriteFile(path, na.Favicon)
	if err != nil {
		return err
	}

	// Write static files
	for p, b := range na.StaticFiles {
		path := filepath.Join(dir, "web/public", p)
		err = WriteFile(path, b)
		if err != nil {
			return err
		}
	}

	// Write src/types
	for id, t := range na.Types {
		path := filepath.Join(dir, "web/src/types", id+".ts")
		err = t.WriteTypeScriptFile(path)
		if err != nil {
			return err
		}
	}

	// Write src/constants
	for id, v := range na.Constants {
		path := filepath.Join(dir, "web/src/constants", id+".ts")
		err = v.WriteTypeScriptConstantFile(path)
		if err != nil {
			return err
		}
	}

	// Write src/functions
	for id, f := range na.Functions {
		path := filepath.Join(dir, "web/src/functions", id+".ts")
		err = f.WriteTypeScriptFile(path)
		if err != nil {
			return err
		}
	}

	// Write src/hooks
	for id, h := range na.Hooks {
		path := filepath.Join(dir, "web/src/hooks", id+".js")
		err = h.WriteTypeScriptFile(path)
		if err != nil {
			return err
		}
	}

	// Write src/components
	for id, c := range na.Components {
		path := filepath.Join(dir, "web/src/components", id+".tsx")
		err = c.Write(path)
		if err != nil {
			return err
		}
	}

	// Write pages
	for id, p := range na.Pages {
		path := filepath.Join(dir, "web/src/app", id, "page.tsx")
		err = p.Write(path)
		if err != nil {
			return err
		}
	}

	// Write layouts
	for id, l := range na.Layouts {
		path := filepath.Join(dir, "web/src/app", id, "layout.tsx")
		err = l.Write(path)
		if err != nil {
			return err
		}
	}

	return nil
}
