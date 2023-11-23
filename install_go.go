package util

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
)

func InstallGo() error {
	osid, _ := GetOSID()
	if osid == "fedora" {
		cmd := exec.Command("dnf", "install", "-y", "golang")
		out, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Println(string(out))
			return err
		}
		return nil
	}

	// Check [1] for latest.
	url := "https://go.dev/dl/go1.21.4.linux-amd64.tar.gz"

	// Download release.
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		return fmt.Errorf("unexpected status code (%d) from go.dev/dl", resp.StatusCode)
	}
	dlpath := filepath.Join(os.TempDir(), "go.tar.gz")
	f, err := os.Create(dlpath)
	if err != nil {
		return err
	}
	_, err = io.Copy(f, resp.Body)
	if err != nil {
		return err
	}

	// Remove existing install
	cmd := exec.Command("rm", "-rf", "/usr/local/go")
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(string(out))
		return err
	}

	// Install
	cmd = exec.Command("tar", "-C", "/usr/local", "-xzf", dlpath)
	out, err = cmd.CombinedOutput()
	if err != nil {
		fmt.Println(string(out))
		return err
	}

	// Add to $PATH
	b, err := os.ReadFile("/etc/profile")
	if err != nil {
		return err
	}
	b = append(b, "export PATH=$PATH:/usr/local/go/bin\n"...)
	err = os.WriteFile("/etc/profile", b, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}
