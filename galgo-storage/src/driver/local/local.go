package local

import (
	"mime/multipart"

	FileHelper "github.com/marceloamoreno87/galgo-storage/src/helper"
)

type StorageLocal struct {
	File *multipart.FileHeader
	Path string
}

func (s *StorageLocal) Get() {
	// get file from local storage

}

func (s *StorageLocal) Save() (err error) {
	err = FileHelper.Save(s.File, s.Path)
	return
}

func (s *StorageLocal) Delete() {
	// delete file from filesystem

}
