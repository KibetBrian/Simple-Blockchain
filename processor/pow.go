package processor

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math"
	"math/big"
)

const targetBits = 44

type ProofOfWork struct {
	Block  *Block
	Target *big.Int
}

func NewProofOfWork(b *Block) *ProofOfWork{
	target := big.NewInt(1)
	target.Lsh(target, 256-targetBits)
	pow := &ProofOfWork{b, target}
	return pow
}

func (pow *ProofOfWork) PrepareData (nonce int) []byte{
	data := bytes.Join([][]byte{
			pow.Block.PrevHash,
		 	pow.Block.Data,
			ConvertToHex(pow.Block.TimeStamp),
			ConvertToHex(int64(targetBits)),
			ConvertToHex(int64(nonce)),
		}, []byte{})
	return data
}

func (pow *ProofOfWork)Run ()(int, []byte){
	var hashint big.Int
	var hash [32]byte
	maxNonce := math.MaxInt64
	nonce:=0

	fmt.Printf("Mining %s block \n",string(pow.Block.Data))
	for nonce < maxNonce {
		data := pow.PrepareData(nonce)
		hash :=sha256.Sum256(data)
		fmt.Printf("\r%x", hash)
		hashint.SetBytes(hash[:])
		if hashint.Cmp(pow.Target)==-1{
			break
		}
		nonce++
	}
	fmt.Print("\n\n")
	return nonce, hash[:]
}

//Validate proof of work 
func(pow *ProofOfWork) ValidatePow () bool{
	var hashint big.Int
	data := pow.PrepareData(pow.Block.Nonce)
	hash := sha256.Sum256(data)
	hashint.SetBytes(hash[:])	
	isValid := hashint.Cmp(pow.Target) == -1
	return isValid
}