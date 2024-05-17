package util

type RemoteFileSystem struct {
	Root string
}

func (fs *RemoteFileSystem) ReadFile(path string) ([]byte, error)
func (fs *RemoteFileSystem) ReadDir(path string) ([]string, error)
func (fs *RemoteFileSystem) IsDir(path string) bool
func (fs *RemoteFileSystem) IsFile(path string) bool
func (fs *RemoteFileSystem) WriteFile(path string, b []byte) error
func (fs *RemoteFileSystem) MakeDir(path string) error
func (fs *RemoteFileSystem) Remove(path string) error
