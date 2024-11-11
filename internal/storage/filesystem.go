package storage

import "os"

type fileSystem interface {
	ReadDir(string) ([]os.DirEntry, error)
	ReadFile(string) ([]byte, error)
}

type osFS struct{}

func (*osFS) ReadDir(path string) ([]os.DirEntry, error) { return os.ReadDir(path) }
func (*osFS) ReadFile(filename string) ([]byte, error)   { return os.ReadFile(filename) }

var storageFS fileSystem = &osFS{}
