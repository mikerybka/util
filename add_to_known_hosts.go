package util

import (
	"errors"
	"fmt"
	"net"
	"os"
	"path/filepath"

	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/agent"
)

func AddToKnownHosts(user, server string) error {
	auth_socket := os.Getenv("SSH_AUTH_SOCK")
	if auth_socket == "" {
		return errors.New("no $SSH_AUTH_SOCK defined")
	}
	conn, err := net.Dial("unix", auth_socket)
	if err != nil {
		return err
	}
	defer conn.Close()
	ag := agent.NewClient(conn)
	auths := []ssh.AuthMethod{ssh.PublicKeysCallback(ag.Signers)}

	config := &ssh.ClientConfig{
		User:            user,
		Auth:            auths,
		HostKeyCallback: keyScanCallback,
	}

	_, err = ssh.Dial("tcp", fmt.Sprintf("%s:%d", server, 22), config)
	if err != nil {
		return err
	}

	return nil
}

func keyScanCallback(hostname string, remote net.Addr, key ssh.PublicKey) error {
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	path := filepath.Join(home, ".ssh/authorized_keys")
	b, _ := os.ReadFile(path)
	b = append(b, []byte(fmt.Sprintf("%s %s\n", hostname[:len(hostname)-3], string(ssh.MarshalAuthorizedKey(key))))...)
	err = os.WriteFile(path, b, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}
