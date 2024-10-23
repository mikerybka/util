package util

import (
	"os"
	"path/filepath"
)

// CopyDir recursively copies a directory from src to dst.
func CopyDir(srcDir, dstDir string) error {
	return filepath.WalkDir(srcDir, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}

		// Determine the destination path
		relPath, err := filepath.Rel(srcDir, path)
		if err != nil {
			return err
		}
		dstPath := filepath.Join(dstDir, relPath)

		if d.IsDir() {
			// Create the directory at the destination path
			return os.MkdirAll(dstPath, os.ModePerm)
		}

		// If it's a file, copy it to the destination
		return CopyFile(path, dstPath)
	})
}
