package gcs

import "mime/multipart"

type StorageGcs struct {
	File *multipart.FileHeader
}

func (s *StorageGcs) Get() {

}

func (s *StorageGcs) Save() (res string, err error) {
	return
}

func (s *StorageGcs) Delete() {

}
