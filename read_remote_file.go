package util

import (
	"bytes"
	"fmt"

	"golang.org/x/crypto/ssh"
)

func ReadRemoteFile(user, password, host string, port int, remotePath string) (string, error) {
	// SSH config
	config := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(), // Use with caution; ideally verify host key
	}

	// Connect
	addr := fmt.Sprintf("%s:%d", host, port)
	client, err := ssh.Dial("tcp", addr, config)
	if err != nil {
		return "", fmt.Errorf("failed to dial: %w", err)
	}
	defer client.Close()

	// Start new session
	session, err := client.NewSession()
	if err != nil {
		return "", fmt.Errorf("failed to create session: %w", err)
	}
	defer session.Close()

	// Prepare output buffer
	var stdout bytes.Buffer
	session.Stdout = &stdout

	// Run command to cat the file
	cmd := fmt.Sprintf("cat %q", remotePath)
	if err := session.Run(cmd); err != nil {
		return "", fmt.Errorf("failed to run remote command: %w", err)
	}

	return stdout.String(), nil
}
