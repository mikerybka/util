package util

import (
	"os"
	"os/exec"
)

type GoBuilder struct {
	Workdir string
}

func (builder *GoBuilder) Build(path string, outfile string) error {
	cmd := exec.Command("go", "build", "-o", outfile, path)
	cmd.Dir = builder.Workdir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
