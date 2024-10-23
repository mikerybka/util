package util

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

func DownloadDir(user, host, path string) ([]byte, error) {
	// Create archive of the path
	remoteTempFile := fmt.Sprintf("/tmp/dl-%s", strconv.Itoa(int(time.Now().Unix())))
	out, err := ExecRemote(user, host, fmt.Sprintf("tar -czvf %s %s", remoteTempFile, path))
	if err != nil {
		return nil, fmt.Errorf("%s: %s", err, out)
	}

	// Donwnload the archive to a local temp file
	localTempFile := filepath.Join(os.TempDir(), "dl-"+strconv.Itoa(int(time.Now().Unix())))
	out, err = CopyFileFromRemote(user, host, remoteTempFile, localTempFile)
	if err != nil {
		return nil, fmt.Errorf("%s: %s", err, out)
	}

	// Delete the remote temp file
	out, err = ExecRemote(user, host, fmt.Sprintf("rm %s", remoteTempFile))
	if err != nil {
		return nil, fmt.Errorf("%s: %s", err, out)
	}

	// Read local temp file into memory
	b, err := os.ReadFile(localTempFile)
	if err != nil {
		return nil, err
	}

	// Delete the local temp file
	err = os.Remove(localTempFile)
	if err != nil {
		return nil, err
	}

	// Return the archive data
	return b, nil
}
