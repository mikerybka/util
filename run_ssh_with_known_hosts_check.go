package util

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/knownhosts"
)

func RunSSHCommandWithKnownHostsCheck(host, user, cmd string, auth []ssh.AuthMethod) ([]byte, error) {
	// Set up the SSH client configuration.
	config := &ssh.ClientConfig{
		User:            user,
		Auth:            auth,
		HostKeyCallback: knownHostsCallback(),
	}

	// Connect to the remote server.
	client, err := ssh.Dial("tcp", host, config)
	if err != nil {
		return nil, err
	}
	defer client.Close()

	// Create an SSH session.
	session, err := client.NewSession()
	if err != nil {
		return nil, err
	}
	defer session.Close()

	// Run the provided SSH command.
	out, err := session.CombinedOutput(cmd)
	if err != nil {
		// Check if the error is due to a known_hosts issue.
		if isKnownHostsError(err) {
			// Attempt to add the new host key to known_hosts.
			if err := addToKnownHosts(host); err != nil {
				return nil, err
			}
			// Retry the SSH command.
			return session.CombinedOutput(cmd)
		}
		return nil, err
	}

	return out, nil
}

func isKnownHostsError(err error) bool {
	// Check if the error message contains the known hosts error message.
	return strings.Contains(err.Error(), "knownhosts: key is unknown")
}

func addToKnownHosts(host string) error {
	sshDir := filepath.Join(os.Getenv("HOME"), ".ssh")
	knownHostsFile := filepath.Join(sshDir, "known_hosts")
	publicKeyFile := filepath.Join(sshDir, "id_rsa.pub")
	b, err := os.ReadFile(publicKeyFile)
	if err != nil {
		return err
	}
	key, _, _, _, err := ssh.ParseAuthorizedKey(b)
	if err != nil {
		return err
	}

	// Append the new host key to the known_hosts file.
	if err := appendToKnownHostsFile(knownHostsFile, host, key); err != nil {
		return err
	}

	return nil
}

func appendToKnownHostsFile(knownHostsFile, host string, key ssh.PublicKey) error {
	// Open the known_hosts file in append mode.
	f, err := os.OpenFile(knownHostsFile, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0600)
	if err != nil {
		return err
	}
	defer f.Close()

	// Format the host key entry and write it to the file.
	_, err = fmt.Fprintf(f, "%s %s\n", host, ssh.MarshalAuthorizedKey(key))
	if err != nil {
		return err
	}

	return nil
}

func knownHostsCallback() ssh.HostKeyCallback {
	knownHostsFile := filepath.Join(os.Getenv("HOME"), ".ssh", "known_hosts")
	if !Exists(knownHostsFile) {
		Touch(knownHostsFile)
	}
	knownHostsCallback, err := knownhosts.New(knownHostsFile)
	if err != nil {
		fmt.Printf("Error loading known_hosts file: %v\n", err)
		knownHostsCallback = ssh.InsecureIgnoreHostKey()
	}
	return knownHostsCallback
}
