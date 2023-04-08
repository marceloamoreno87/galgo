package storage

type StorageInterface interface {
	Save() (err error)
	Delete(filename string) (err error)
}

type StorageExternalInterface interface {
	Get(filename string) (err error)
	StorageInterface
}
