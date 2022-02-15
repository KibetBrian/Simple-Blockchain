package processor

import (
	"blockchain/structs"
	"crypto/sha256"
	"time"
)

func GenerateNewHash(prevHash []byte, transactions []string, time time.Time) []byte{
	input := append(prevHash, time.String()...)

	for transaction :=range transactions{
		input=append(input, string(rune(transaction))...)
	}
	hash := sha256.Sum256(input)
	return hash[:]
}

func CreateNewBlock(prevHash []byte, transactions []string) *structs.Block{
	currentTime := time.Now()
	block :=&structs.Block{
		PrevHash: prevHash,
		Transactions: transactions,
		TimeStamp: currentTime,
		Hash: GenerateNewHash(prevHash,transactions,currentTime),
	}
	return block
}

