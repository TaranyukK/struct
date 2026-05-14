package storage

import (
	"encoding/json"
	"fmt"
	"struct/bins"
	"struct/file"
)

func StorageBins(bins []*bins.Bin) {
	serializedBins, err := json.Marshal(bins)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	file.WriteFile(serializedBins, "bins.json")
}

func ReadBins(name string) []*bins.Bin {
	data, err := file.ReadFile(name)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	var bins []*bins.Bin
	err = json.Unmarshal(data, &bins)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	return bins
}
