package util

import (
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

}
