package bins

import (
	"errors"
	"struct/storage"
	"time"
)

type Bin struct {
	Id        string    `json:"id"`
	Private   bool      `json:"private"`
	CreatedAt time.Time `json:"createdAt"`
	Name      string    `json:"name"`
}

var BinList []*Bin

func NewBin(id string, private bool) (*Bin, error) {
	if id == "" {
		return nil, errors.New("INVALID_ID")
	}
	newBin := &Bin{
		Id:        id,
		Private:   private,
		CreatedAt: time.Now(),
	}
	BinList = append(BinList, newBin)
	storage.StorageBins(BinList)
	return newBin, nil
}
