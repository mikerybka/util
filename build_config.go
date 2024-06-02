package util

import (
	"fmt"
	"os/exec"
)

type BuildConfig struct {
	Type string
}

func (config *BuildConfig) Build(path string, out string) error {
	switch config.Type {
	case "go":
		cmd := exec.Command("go", "build", "-o", out, ".")
		cmd.Dir = path
		out, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Println(string(out))
			return err
		}
		return nil
	default:
		panic(fmt.Sprintf("unknown type: %s", config.Type))
	}
}
