package util

import (
	"crypto/sha256"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/fsnotify/fsnotify"
)

// computeFileHash returns the SHA-256 hash of the file's contents.
func computeFileHash(path string) ([]byte, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	hasher := sha256.New()
	if _, err := io.Copy(hasher, file); err != nil {
		return nil, err
	}
	return hasher.Sum(nil), nil
}

// OnFileChange sets up a watcher that calls `callback()` ONCE each time the file's contents change.
func OnFileChange(path string, callback func()) error {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return fmt.Errorf("failed to create watcher: %w", err)
	}

	err = watcher.Add(path)
	if err != nil {
		return fmt.Errorf("failed to watch file %s: %w", path, err)
	}

	// Track last known hash
	lastHash, err := computeFileHash(path)
	if err != nil {
		log.Printf("initial hash error: %v", err)
		lastHash = nil // Treat as unknown
	}

	go func() {
		defer watcher.Close()
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}

				if event.Op&(fsnotify.Write|fsnotify.Create) != 0 {
					currentHash, err := computeFileHash(path)
					if err != nil {
						log.Printf("hash error: %v", err)
						continue
					}

					if !hashEqual(lastHash, currentHash) {
						lastHash = currentHash
						go callback()
					}
				}

			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Printf("watcher error: %v", err)
			}
		}
	}()
	return nil
}

// hashEqual compares two SHA-256 hashes.
func hashEqual(a, b []byte) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
