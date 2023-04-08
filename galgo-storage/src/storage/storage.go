package storage

func Get(s StorageInterface, filename string) (err error) {
	err = s.Get(filename)
	return
}

func Save(s StorageInterface) (err error) {
	err = s.Save()
	return
}

func Delete(s StorageInterface, filename string) (err error) {
	err = s.Delete(filename)
	return
}
