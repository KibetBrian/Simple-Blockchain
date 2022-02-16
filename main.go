package main

import (
	"blockchain/processor"
)

func main() {
	chain := processor.NewBlockChain();
	chain.AddBlock("Just after genesis")
	processor.PrintBlockChain(chain)
}