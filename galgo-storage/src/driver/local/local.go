package local

import (
	"mime/multipart"

	File "github.com/marceloamoreno87/galgo-storage/src/file"
)

type StorageLocal struct {
	File *multipart.FileHeader
	Path string
}

func (s *StorageLocal) Get(path string, filename string) (err error) {
	err = File.Get(path, filename)
	return
}

func (s *StorageLocal) Save() (err error) {
	err = File.Save(s.File, s.Path)
	return
}

func (s *StorageLocal) Delete(path string, filename string) (err error) {
	err = File.Delete(path, filename)
	return
}
