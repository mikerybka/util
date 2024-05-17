package util

type RemoteFileSystem struct {
	Root string
}

func (fs *RemoteFileSystem) ReadFile(path string) ([]byte, error) {
	panic("not implemented")
}
func (fs *RemoteFileSystem) ReadDir(path string) ([]string, error) {
	panic("not implemented")
}
func (fs *RemoteFileSystem) IsDir(path string) bool {
	panic("not implemented")
}
func (fs *RemoteFileSystem) IsFile(path string) bool {
	panic("not implemented")
}
func (fs *RemoteFileSystem) WriteFile(path string, b []byte) error {
	panic("not implemented")
}
func (fs *RemoteFileSystem) MakeDir(path string) error {
	panic("not implemented")
}
func (fs *RemoteFileSystem) Remove(path string) error {
	panic("not implemented")
}
