package util

import (
	"encoding/json"
	"net/http"
	"os/exec"
	"path/filepath"
)

type CmdServer struct {
	Dir string
}

func (s *CmdServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	req := &struct {
		Dir string
		Cmd string
	}{}
	json.NewDecoder(r.Body).Decode(req)
	cmd := exec.Command("/bin/sh", "-C", req.Cmd)
	cmd.Dir = filepath.Join(s.Dir, req.Dir)
	out, err := cmd.CombinedOutput()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	res := &struct {
		Output   string
		ExitCode int
	}{
		Output:   string(out),
		ExitCode: cmd.ProcessState.ExitCode(),
	}
	json.NewEncoder(w).Encode(res)
}
