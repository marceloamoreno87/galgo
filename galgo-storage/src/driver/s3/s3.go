package s3

import "mime/multipart"

type StorageS3 struct {
	File *multipart.FileHeader
}

func (s *StorageS3) Get() {

}

func (s *StorageS3) Save() (res string, err error) {
	return
}

func (s *StorageS3) Delete() {

}
