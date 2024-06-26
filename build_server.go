package util

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httputil"
	"path/filepath"
	"strings"
)

type BuildServer struct {
	Workdir string
	Config  map[string]*BuildConfig
}

func (s *BuildServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	b, err := httputil.DumpRequest(r, true)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))

	if strings.HasPrefix(r.URL.Path, "/builds") {
		buildDir := filepath.Join(s.Workdir, "builds")
		http.StripPrefix("/builds", http.FileServer(http.Dir(buildDir))).ServeHTTP(w, r)
		return
	}

	if r.URL.Path == "/webhooks/github" {
		var req GithubWebhookRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			return
		}

		path := filepath.Join(s.Workdir, "src", req.Repository.FullName)
		err = req.Repository.SyncLocal(path)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		return
	}
}
