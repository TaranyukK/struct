package main

import (
	"fmt"
	"struct/bins"
)

func main() {
	id := "1"
	private := false
	newBin, err := bins.NewBin(id, private)
	if err != nil {
		if err.Error() == "INVALID_ID" {
			fmt.Println("Неверный id")
		}
		return
	}
	fmt.Println(newBin)
}
