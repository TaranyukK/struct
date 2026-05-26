package main

import (
	"fmt"
	bin "struct/bins"
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
	existingBins := storage.Read("bins.json")
	existingBins = append(existingBins, *newBin)
	storage.Write(existingBins)
}
