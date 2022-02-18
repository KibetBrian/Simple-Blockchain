package processor

import (
	"bytes"
	"crypto/sha256"
	"time"
)

type Block struct {
	PrevHash  []byte
	BlockHash []byte
	TimeStamp int64
	Data      []byte
	Nonce	  int
}

//Function to create new block
func GenerateNewBlock(prevHash []byte, data string) *Block {
	currentTime := time.Now().UnixNano()
	block := &Block{
		PrevHash:  prevHash,
		Data:      []byte(data),
		TimeStamp: currentTime,
	}
	pow := NewProofOfWork(block)
	nonce, hash := pow.Run()
	block.BlockHash =hash
	block.Nonce = nonce
	return block
}

//Function to get block hash from given metadata
func GenerateNewHash(prevHash []byte, data string, time uint64) []byte{
	timeBytes := ConvertToBytes(time)
	metaData:= bytes.Join([][]byte{prevHash,[]byte(data), timeBytes}, []byte{})
	metaDataHash := sha256.Sum256(metaData)
	return metaDataHash[:]
}


