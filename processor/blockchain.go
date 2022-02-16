package processor

type Blockchain struct{
	Blocks []*Block
}

//Function to add new block to create chain
func (b *Blockchain) AddBlock (data string){
	prevBlockHash :=  b.Blocks[len(b.Blocks)-1].BlockHash
	newBlock := GenerateNewBlock(prevBlockHash, data)
	b.Blocks = append(b.Blocks, newBlock)
}

//Creates the first block
func CreateGenesisBlock () *Block{
	genesisBlock :=GenerateNewBlock([]byte{}, "Genesis")
	return genesisBlock
}

//Returns new blockchain
func NewBlockChain () *Blockchain{
	chain := &Blockchain{[]*Block{CreateGenesisBlock()}}
	return chain  
}
