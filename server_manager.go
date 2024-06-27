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
	service := &SystemdService{
		Name:  "server",
		Desc:  "server",
		After: "network.target",
		Type:  "simple",
		Env: []Pair[string, string]{
			{
				K: "TWILIO_ACCOUNT_SID",
				V: config.TwilioClient.AccountSID,
			},
			{
				K: "TWILIO_AUTH_TOKEN",
				V: config.TwilioClient.AuthToken,
			},
			{
				K: "TWILIO_PHONE_NUMBER",
				V: config.TwilioClient.PhoneNumber,
			},
			{
				K: "DATA_DIR",
				V: config.DataDir,
			},
			{
				K: "ADMIN_PHONE",
				V: config.AdminPhone,
			},
			{
				K: "ADMIN_EMAIL",
				V: config.AdminEmail,
			},
			{
				K: "CERT_DIR",
				V: config.CertDir,
			},
		},
		AutoRestart: "on-failure",
		WantedBy:    "multi-user.target",
	}
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
