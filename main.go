package main

import (
	"blockchain/processor"
	"blockchain/utils"
)

func main() {
	genesisTransaction := []string{"Initial Transaction"}
	genesisBlock := processor.CreateNewBlock([]byte{}, genesisTransaction)
	utils.PrintBlockInformation(genesisBlock)
}