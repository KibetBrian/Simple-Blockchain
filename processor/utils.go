package processor

import (
	"encoding/binary"
	"fmt"
)

//Function to convert time in unix nano to slice of bytes
func ConvertToBytes (time uint64) []byte{
	timeStampBytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(timeStampBytes, uint64(time))
	return timeStampBytes
}

//Function to print block information
func PrintBlock (block *Block){
	fmt.Println("PrevHash: ", block.PrevHash)
	fmt.Println("Time ", block.TimeStamp)
	fmt.Printf("Hash: %x", block.BlockHash)
	printData(block)
}

//Function to get data in the data slices on a block
func printData(block *Block){
	for _, v:= range block.Data{
		fmt.Println("\t" ,string(v))
	}
}

func PrintBlockChain (blockChain *Blockchain){
	for i:=0; i<len(blockChain.Blocks);i++{
		fmt.Printf("Block %v: Hash %x \n", i,string(blockChain.Blocks[i].BlockHash))
	}
}

//Takes type int64 as input and return bytes
func ConvertToHex (input int64) []byte {
	 b := make([]byte, 8)
	 binary.LittleEndian.PutUint64(b, uint64(input))
	return b
}
