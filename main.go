package main

import (
	"fmt"
	bin "struct/bins"
	"struct/file"
	"struct/storage"
)

func main() {
	id := "1"
	private := false
	name := "Some File"
	newBin, err := bin.NewBin(id, private, name)
	if err != nil {
		fmt.Println("Ошибка создания:", err)
		return
	}
	osFile := file.NewOSFile()
	jsonStorage := storage.NewJSONBinStorage(osFile)
	existingBins := jsonStorage.Read("bins.json")
	existingBins = append(existingBins, *newBin)
	jsonStorage.Write(existingBins)
}
