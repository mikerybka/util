package util

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
)

type GithubMirror struct {
	Workdir string
}

// Handle webhooks from github
func (g *GithubMirror) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		return
	}
	req := ReadJSON[*GithubWebhookRequest](r.Body)
	err := g.Pull(req.Repository.FullName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (g *GithubMirror) Pull(repo string) error {
	path := filepath.Join(g.Workdir, repo)
	if !IsDir(path) {
		err := os.MkdirAll(filepath.Dir(path), os.ModePerm)
		if err != nil {
			panic(err)
		}
		cmd := exec.Command("gh", "repo", "clone")
		cmd.Dir = filepath.Dir(path)
		out, err := cmd.CombinedOutput()
		if err != nil {
			return fmt.Errorf("%s: %s", err, out)
		}
		return nil
	}
	r := &GitRepo{
		Dir: path,
	}
	_, err := r.Pull()
	return err
}
