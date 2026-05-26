package storage

import (
	"encoding/json"
	"fmt"
	bin "struct/bins"
	"struct/file"
)

type Bins interface {
	Read(string) []bin.Bin
	Write([]bin.Bin)
}

type JSONBinStorage struct {
	FileRepo file.File
}

func NewJSONBinStorage(file file.File) *JSONBinStorage {
	return &JSONBinStorage{
		FileRepo: file,
	}
}

func (storage *JSONBinStorage) Write(bins []bin.Bin) {
	serializedBins, err := json.Marshal(bins)
	if err != nil {
		fmt.Println("Ошибка сериализации:", err)
		return
	}
	storage.FileRepo.Write(serializedBins, "bins.json")
}

func (storage *JSONBinStorage) Read(name string) []bin.Bin {
	data, err := storage.FileRepo.Read(name)
	if err != nil {
		return nil
	}
	var bins []bin.Bin
	err = json.Unmarshal(data, &bins)
	if err != nil {
		fmt.Println("Ошибка десериализации:", err)
		return nil
	}
	return bins
}
