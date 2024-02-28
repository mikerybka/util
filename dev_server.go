package util

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"
)

type DevServer struct {
	PkgPath string
	cmd     *exec.Cmd
	mutex   sync.Mutex
}

func (s *DevServer) watchDir() string {
	parts := ParsePath(s.PkgPath)
	if len(parts) >= 3 {
		parts = parts[:3]
	}
	return filepath.Join(HomeDir(), "src", strings.Join(parts, "/"))
}

func (s *DevServer) Start() {
	s.build()
	s.run()
	s.watch()
}

func (s *DevServer) build() {
	s.cmd = exec.Command("go", "build", "-o", s.exePath(), s.PkgPath)
	s.cmd.Stderr = os.Stderr
	s.cmd.Stdout = os.Stdout
	s.cmd.Run()
}

func (s *DevServer) exePath() string {
	return filepath.Join(os.TempDir(), filepath.Base(s.PkgPath))
}

func (s *DevServer) run() {
	s.cmd = exec.Command(s.exePath())
	s.cmd.Stderr = os.Stderr
	s.cmd.Stdout = os.Stdout
	s.cmd.Start()
}

func (s *DevServer) kill() {
	s.cmd.Process.Kill()
}

func (s *DevServer) watch() {
	Watch(s.watchDir(), func() {
		s.mutex.Lock()
		defer s.mutex.Unlock()
		s.build()
		s.kill()
		s.run()
	})
}
