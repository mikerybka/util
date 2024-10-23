package util

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

func UploadDir(user, host string, b []byte, path string) error {
	// Write archive to local temp file
	localTempFile := filepath.Join(os.TempDir(), "ul-"+strconv.Itoa(int(time.Now().Unix())))
	err := os.WriteFile(localTempFile, b, os.ModePerm)
	if err != nil {
		return err
	}

	// Copy local temp file to server
	remoteTempFile := filepath.Join("/tmp", "ul-"+strconv.Itoa(int(time.Now().Unix())))
	out, err := CopyFileToRemote(user, host, localTempFile, remoteTempFile)
	if err != nil {
		return fmt.Errorf("%s: %s", err, out)
	}

	// Delete local temp file
	err = os.Remove(localTempFile)
	if err != nil {
		return err
	}

	// Unpack archive on server
	out, err = ExecRemote(user, host, fmt.Sprintf("tar -xzvf %s -C %s", remoteTempFile, path))
	if err != nil {
		return fmt.Errorf("%s: %s", err, out)
	}

	// Delete archive file on server
	out, err = ExecRemote(user, host, fmt.Sprintf("rm %s", remoteTempFile))
	if err != nil {
		return fmt.Errorf("%s: %s", err, out)
	}

	return nil
}
