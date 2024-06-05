package util

import (
	"fmt"
	"os/exec"
)

type GoCodebase struct {
	Dir string
}

func (c *GoCodebase) UpdateDeps() (bool, error) {
	cmd := exec.Command("go", "get", "-u")
	cmd.Dir = c.Dir
	out, err := cmd.CombinedOutput()
	if err != nil {
		if len(out) > 0 {
			return false, fmt.Errorf("%s: %s", err, out)
		}
		return false, err
	}
	if len(out) > 0 {
		return true, nil
	}

	// Check if there are changes.
	cmd = exec.Command("git", "diff", "--exit-code")
	cmd.Dir = c.Dir
	out, err = cmd.CombinedOutput()
	if err != nil {
		if len(out) > 0 {
			return false, fmt.Errorf("%s: %s", err, out)
		}
		return false, err
	}

	return len(out) > 0, nil
}

func (c *GoCodebase) GitRepo() *GitRepo {
	return &GitRepo{
		Dir: c.Dir,
	}
}

func (c GoCodebase) Build(outFile string) error {
	cmd := exec.Command("go", "build", "-o", outFile, ".")
	cmd.Dir = c.Dir
	out, err := cmd.CombinedOutput()
	if err != nil {
		if len(out) > 0 {
			return fmt.Errorf("%s: %s", err, out)
		}
		return err
	}
	if len(out) > 0 {
		return fmt.Errorf("go build: %s", out)
	}
	return nil
}
