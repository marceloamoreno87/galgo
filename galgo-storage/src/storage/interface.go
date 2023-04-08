package storage

type StorageInterface interface {
	Get(filename string) (err error)
	Save() (err error)
	Delete(filename string) (err error)
}
