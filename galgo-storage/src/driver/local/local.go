package local

import (
	"mime/multipart"

	File "github.com/marceloamoreno87/galgo-storage/src/file"
)

type StorageLocal struct {
	File *multipart.FileHeader
}

func (s *StorageLocal) Get(filename string) (err error) {
	err = File.Get(filename)
	return
}

func (s *StorageLocal) Save() (err error) {
	err = File.Save(s.File)
	return
}

func (s *StorageLocal) Delete(filename string) (err error) {
	err = File.Delete(filename)
	return
}
