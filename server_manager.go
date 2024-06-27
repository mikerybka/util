package util

import (
	"path/filepath"
)

type ServerManager struct {
	RootDir string
}

func (m *ServerManager) SetAppConfig(host string, config *AppConfig) error {
	path := filepath.Join(m.RootDir, host, "config.json")
	return WriteJSONFile(path, config)
}

func (m *ServerManager) SetServerConfig(host string, config *Server) error
