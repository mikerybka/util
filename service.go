package util

import "fmt"

type Service struct {
	Name  string
	Env   map[string]string
	Start string
	User  string
}

func (s *Service) Systemd() string {
	env := ""
	for k, v := range s.Env {
		env += fmt.Sprintf("Environment=%s=%s\n", k, v)
	}
	return fmt.Sprintf(`[Unit]
Description=%s
After=network.target

[Service]
Type=simple
%sExecStart=%s
Restart=on-failure
User=%s

[Install]
WantedBy=multi-user.target`, s.Name, env, s.Start, s.User)
}
