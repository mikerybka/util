package util

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"text/template"
)

type SystemdService struct {
	Name        string
	Desc        string
	After       string
	Type        string
	Env         []Pair[string, string]
	ExecStart   string
	AutoRestart string
	WantedBy    string
}

func (s *SystemdService) WriteConfigToFile() error {
	path := fmt.Sprintf("/etc/systemd/system/%s.service", s.Name)
	f, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("create: %s", err)
	}
	return s.WriteConfig(f)
}

func (s *SystemdService) WriteConfig(w io.Writer) error {
	return template.Must(template.New("service").Parse(`[Unit]
	Description={{ .Desc }}
	After={{ .After }}
	
	[Service]
	Type={{ .Type }}
	{{ range .Env }}
	Environment="{{ .K }}={{ .V }}"
	{{ end }}
	ExecStart={{ .ExecStart }}
	Restart={{ .AutoRestart }}
	
	[Install]
	WantedBy={{ .WantedBy }}
	`)).Execute(w, s)
}

func (s *SystemdService) Restart() error {
	cmd := exec.Command("systemctl", "restart", s.Name)
	out, err := cmd.CombinedOutput()
	if err != nil {
		if len(out) > 0 {
			return fmt.Errorf("%s: %s", err, out)
		}
		return err
	}
	if len(out) > 0 {
		return fmt.Errorf("systemctl restart %s: %s", s.Name, out)
	}
	return nil
}
