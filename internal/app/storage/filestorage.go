package storage

import (
	"io"
	"os"
	"path/filepath"
)

type FileStorage interface {
	Exists(name string) (bool, error)
	Save(name string) (io.WriteCloser, error)
	Get(name string) (io.ReadCloser, error)
}

type SimpleFileStorage struct {
	basePath string
}

func NewSimpleFileStorage(basePath string) FileStorage {
	return &SimpleFileStorage{
		basePath: basePath,
	}
}

func (fs *SimpleFileStorage) Exists(name string) (bool, error) {
	return true, nil
}

func (fs *SimpleFileStorage) Save(name string) (io.WriteCloser, error) {
	path := filepath.Join(fs.basePath, name)
	file, err := os.Create(path)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func (fs *SimpleFileStorage) Get(name string) (io.ReadCloser, error) {
	path := filepath.Join(fs.basePath, name)
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	return file, nil
}
