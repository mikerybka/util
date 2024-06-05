package util

import (
	"fmt"
	"os/exec"
)

type GitRepo struct {
	Dir string
}

// Pull from the default upstream.
// Returns true if there were updates, if git prints "Already up to date.", false is returned.
// Other errors are returned.
func (r *GitRepo) Pull() (bool, error) {
	cmd := exec.Command("git", "pull")
	cmd.Dir = r.Dir
	out, err := cmd.CombinedOutput()
	if err != nil && len(out) > 0 {
		return false, fmt.Errorf("%s: %s", err, out)
	}

	if string(out) == "Already up to date.\n" {
		return false, nil
	} else {
		return true, nil
	}
}

func (r *GitRepo) Push() error {
	cmd := exec.Command("git", "push")
	cmd.Dir = r.Dir
	out, err := cmd.CombinedOutput()
	if err != nil && len(out) > 0 {
		return fmt.Errorf("%s: %s", err, out)
	}
	return err
}

func (r *GitRepo) AddAll() error {
	cmd := exec.Command("git", "add", "--all")
	cmd.Dir = r.Dir
	out, err := cmd.CombinedOutput()
	if err != nil && len(out) > 0 {
		return fmt.Errorf("%s: %s", err, out)
	}
	return err
}

func (r *GitRepo) Commit(msg string) error {
	cmd := exec.Command("git", "commit", "-m", msg)
	cmd.Dir = r.Dir
	out, err := cmd.CombinedOutput()
	if err != nil && len(out) > 0 {
		return fmt.Errorf("%s: %s", err, out)
	}
	return err
}
