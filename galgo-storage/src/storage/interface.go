package storage

type StorageInterface interface {
	Get(path string, filename string) (err error)
	Save() (err error)
	Delete(path string, filename string) (err error)
}
