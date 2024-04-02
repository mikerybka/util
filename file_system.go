package util

import (
	"fmt"

	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

type FileSystem struct {
	HetznerStorageBoxAddr     string
	HetznerStorageBoxUsername string
	HetznerStorageBoxPassword string
}

func (fs *FileSystem) NewSSHClient() (*ssh.Client, error) {
	var hostKey ssh.PublicKey
	// An SSH client is represented with a ClientConn.
	//
	// To authenticate with the remote server you must pass at least one
	// implementation of AuthMethod via the Auth field in ClientConfig,
	// and provide a HostKeyCallback.
	config := &ssh.ClientConfig{
		User: fs.HetznerStorageBoxUsername,
		Auth: []ssh.AuthMethod{
			ssh.Password(fs.HetznerStorageBoxPassword),
		},
		HostKeyCallback: ssh.FixedHostKey(hostKey),
	}
	client, err := ssh.Dial("tcp", fmt.Sprintf("%s:23", fs.HetznerStorageBoxAddr), config)
	if err != nil {
		return nil, fmt.Errorf("Failed to dial: %s", err)
	}
	return client, nil
}

func (fs *FileSystem) NewSFTPClient() (*sftp.Client, error) {
	conn, err := fs.NewSSHClient()
	if err != nil {
		return nil, err
	}
	return sftp.NewClient(conn)
}
