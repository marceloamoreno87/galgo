package storage

func Get(s StorageInterface, path string, filename string) (err error) {
	err = s.Get(path, filename)
	return
}

func Save(s StorageInterface) (err error) {
	err = s.Save()
	return
}

func Delete(s StorageInterface, path string, filename string) (err error) {
	err = s.Delete(path, filename)
	return
}
