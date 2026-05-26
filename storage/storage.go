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

func Write(bins []bin.Bin) {
	serializedBins, err := json.Marshal(bins)
	if err != nil {
		fmt.Println("Ошибка сериализации:", err)
		return
	}
	file.Write(serializedBins, "bins.json")
}

func Read(name string) []bin.Bin {
	data, err := file.Read(name)
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
