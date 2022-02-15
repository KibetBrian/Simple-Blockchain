package utils

import (
	"blockchain/structs"
	"fmt"
)

func PrintBlockInformation (block *structs.Block){
	fmt.Println("PrevHash: ", block.PrevHash)
	fmt.Println("Time ", block.TimeStamp)
	fmt.Println("Hash: ", string(block.Hash))
	printTransactions(block)
}

func printTransactions(block *structs.Block){
	for _, v:= range block.Transactions{
		fmt.Println("\t" ,v)
	}
}