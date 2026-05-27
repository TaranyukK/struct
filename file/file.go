package file

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

type File interface {
	Read(string) ([]byte, error)
	Write([]byte, string)
}

type OSFile struct{}

func NewOSFile() *OSFile {
	return &OSFile{}
}

func (file *OSFile) Read(name string) ([]byte, error) {
	if filepath.Ext(name) != ".json" {
		fmt.Println("Недопустимый формат файла")
		err := errors.New("INVALID_FORMAT")
		return nil, err
	}
	data, err := os.ReadFile(name)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (file *OSFile) Write(content []byte, name string) {
	osfile, err := os.Create(name)
	if err != nil {
		fmt.Println(err)
	}
	defer osfile.Close()
	_, err = osfile.Write(content)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Запись успешна")
}
