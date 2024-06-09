package util

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type GoWorkspace struct {
	Dir string
}

func (w *GoWorkspace) Build(pkg string, o string) error {
	cmd := exec.Command("go", "build", "-o", o, pkg)
	cmd.Dir = w.Dir
	out, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("go build: %s: %s", err.Error(), out)
	}
	return nil
}

func (w *GoWorkspace) Clone(pkg string) error {
	if !strings.HasPrefix(pkg, "github.com/") {
		return fmt.Errorf("only github packages are supported")
	}

	// Parse the ghID.
	ghID := strings.TrimPrefix(pkg, "github.com/")
	org, repo, found := strings.Cut(ghID, "/")
	if !found {
		panic("expected full gh repo id")
	}

	// Make sure the org dir exists.
	orgPath := filepath.Join(w.Dir, "src/github.com", org)
	err := os.MkdirAll(orgPath, os.ModePerm)
	if err != nil {
		return err
	}

	// Clone the repo using gh.
	cmd := exec.Command("gh", "repo", "clone", ghID)
	cmd.Dir = orgPath
	out, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("gh repo clone: %s: %s", err.Error(), out)
	}

	// Add the repo to the workspace.
	repoPath := filepath.Join(orgPath, repo)
	cmd = exec.Command("go", "work", "use", ".")
	cmd.Dir = repoPath
	out, err = cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("go work use: %s: %s", err.Error(), out)
	}

	return nil
}

func (w *GoWorkspace) Pull(pkg string) error {
	cmd := exec.Command("git", "pull")
	cmd.Dir = filepath.Join(w.Dir, "src", pkg)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("gh repo clone: %s: %s", err.Error(), out)
	}
	return nil
}
