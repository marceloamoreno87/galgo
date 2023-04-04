package helper

import (
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
)

func Save(file *multipart.FileHeader, path string) (err error) {
	src, err := file.Open()
	if err != nil {
		return
	}
	defer src.Close()
	if err = os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return
	}

	out, err := os.Create(path + file.Filename)
	if err != nil {
		return
	}
	defer out.Close()

	_, err = io.Copy(out, src)
	return
}

func Delete(filename string) error {
	return os.Remove(filename)
}
