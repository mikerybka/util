package util

import (
	"fmt"
	"os/exec"
)

type BuildConfig struct {
	Type      string
	Path      string
	Out       string
	OnSuccess string
}

func (config *BuildConfig) Build() error {
	switch config.Type {
	case "go":
		cmd := exec.Command("go", "build", "-o", config.Out, ".")
		cmd.Dir = config.Path
		out, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Println(string(out))
			return err
		}
	default:
		panic(fmt.Sprintf("unknown type: %s", config.Type))
	}

	// Run post-build script
	cmd := exec.Command("bash", "-c", config.OnSuccess)
	cmd.Dir = config.Path
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(string(out))
		return err
	}

	return nil
}
