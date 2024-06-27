package util

import (
	"fmt"
	"path/filepath"
)

type ServerManager struct{}

func (m *ServerManager) SetAppConfig(host string, config *AppConfig) error {
	path := filepath.Join("/root", host, "config.json")
	return WriteJSONFile(path, config)
}

func (m *ServerManager) SetServerConfig(config *Server) error {
	service := config.SystemdService()
	err := service.WriteConfigToFile()
	if err != nil {
		return fmt.Errorf("writing systemd service: %s", err)
	}
	err = ReloadSystemd()
	if err != nil {
		return fmt.Errorf("reloading systemd: %s", err)
	}
	err = service.Restart()
	if err != nil {
		return fmt.Errorf("restarting server: %s", err)
	}
	return nil
}
