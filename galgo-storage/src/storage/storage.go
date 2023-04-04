package storage

func Get(s StorageInterface) (err error) {
	return

}

func Save(s StorageInterface) (err error) {
	err = s.Save()
	return
}

func Delete(s StorageInterface) (err error) {
	return
}
