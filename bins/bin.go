package bins

import (
	"errors"
	"time"
)

type Bin struct {
	id        string
	private   bool
	createdAt time.Time
	name      string
}

var BinList []*Bin

func NewBin(id string, private bool) (*Bin, error) {
	if id == "" {
		return nil, errors.New("INVALID_ID")
	}

	newBin := &Bin{
		id:        id,
		private:   private,
		createdAt: time.Now(),
	}

	BinList = append(BinList, newBin)

	return newBin, nil
}
