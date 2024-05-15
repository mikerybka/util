package util

// FileStore stores and serves files.
// To use, set up a directory with an `auth.json` file and a `content` directory.
type FileStore struct {
	Workdir string
}

func (fs *FileStore) HasAccess(userID, path string) bool
func (fs *FileStore) ReadFile(path string) ([]byte, error)
