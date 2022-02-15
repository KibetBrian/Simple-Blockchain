package structs

import (
	"time"
)

type Block struct {
	PrevHash     []byte
	Transactions []string
	Hash         []byte
	TimeStamp    time.Time
}

type Blockchain struct{
	Blocks []*Block
}
