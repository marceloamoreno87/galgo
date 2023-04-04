package azs

import "mime/multipart"

type StorageAzs struct {
	File *multipart.FileHeader
}

func (s *StorageAzs) Get() {

}

func (s *StorageAzs) Save() (res string, err error) {
	return
}

func (s *StorageAzs) Delete() {

}
