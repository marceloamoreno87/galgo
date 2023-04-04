package storage

type StorageInterface interface {
	Get()
	Save() (err error)
	Delete()
}
