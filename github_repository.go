package util

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

type GithubRepository struct {
	ID       int    `json:"id"`
	NodeID   string `json:"node_id"`
	FullName string `json:"full_name"`
}

func (repo *GithubRepository) SyncLocal(path string) error {
	// Check if path exists
	_, err := os.Stat(path)

	// If it's not found, clone the repo
	if errors.Is(err, os.ErrNotExist) {
		dir := filepath.Dir(path)
		err = os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			return err
		}

		cmd := exec.Command("gh", "repo", "clone", repo.FullName)
		cmd.Dir = dir
		out, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Println(string(out))
			return err
		}

		return nil
	}

	// If it is found, try to pull for updates.
	cmd := exec.Command("git", "pull")
	cmd.Dir = path
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(string(out))
		return err
	}

	return nil
}
