package storage

import (
	"encoding/json"
	"fmt"
	"struct/file"
)

func StorageBins(bins []Bin) {
	serializedBins, err := json.Marshal(bins)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	file.WriteFile(serializedBins, "bins.json")
}

func ReadBins(name string) {
	data, err := file.ReadFile(name)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	var bins []Bin
	err = json.Unmarshal(data, &bins)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(bins)
}
