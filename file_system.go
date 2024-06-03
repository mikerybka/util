package util

type FileSystem interface {
	ReadFile(path string) ([]byte, error)
	ReadDir(path string) ([]string, error)
	IsDir(path string) bool
	IsFile(path string) bool
	WriteFile(path string, b []byte) error
	MakeDir(path string) error
	Remove(path string) error
	Dig(path string) FileSystem
}
