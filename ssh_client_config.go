package util

import (
	"os"
	"path/filepath"

	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/knownhosts"
)

func SSHClientConfig(user string) (*ssh.ClientConfig, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		home = "/root"
	}
	path := filepath.Join(home, ".ssh/id_rsa")
	pKey, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	signer, err := ssh.ParsePrivateKey(pKey)
	if err != nil {
		return nil, err
	}

	path = filepath.Join(home, ".ssh/known_hosts")
	hostkeyCallback, err := knownhosts.New(path)
	if err != nil {
		return nil, err
	}

	return &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
		HostKeyCallback: hostkeyCallback,
	}, nil
}
