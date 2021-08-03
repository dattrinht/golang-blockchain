package main

import (
	"fmt"

	"github.com/dattrinht/golang-blockchain/blockchain"
)

func main() {
	chain := blockchain.InitBlockChain()
	chain.AddBlock("1st Block")
	chain.AddBlock("2nd Block")
	chain.AddBlock("3rd Block")

	for _, block := range chain.Blocks {
		fmt.Printf("PrevHash: %x\n", block.PrevHash)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Printf("Data %s\n\n", block.Data)
	}
}
