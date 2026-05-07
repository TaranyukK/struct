package main

import (
	"errors"
	"fmt"
	"time"
)

type Bin struct {
	id        string
	private   bool
	createdAt time.Time
	name      string
}

func NewBin(id string, private bool) (*Bin, error) {
	if id == "" {
		return nil, errors.New("INVALID_ID")
	}

	newBin := &Bin{
		id:        id,
		private:   private,
		createdAt: time.Now(),
	}

	var BinList []*Bin
	BinList = append(BinList, newBin)

	return newBin, nil
}

func main() {
	id := "1"
	private := false

	newBin, err := NewBin(id, private)
	if err != nil {
		if err.Error() == "INVALID_ID" {
			fmt.Println("Неверный id")
		}
		return
	}

	fmt.Println(newBin)
}
