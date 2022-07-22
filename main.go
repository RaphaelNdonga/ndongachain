package main

import (
	"fmt"

	"github.com/RaphaelNdonga/ndongachain/blockchain"
)

func main() {
	chain := blockchain.InitBlockChain()

	chain.AddBlock("First Block after genesis")
	chain.AddBlock("Second Block")
	chain.AddBlock("Third Block")

	for _, block := range chain.Blocks {
		fmt.Printf("Current Hash: %x\n", block.Hash)
		fmt.Printf("Prev Hash: %x\n ", block.PrevHash)
		fmt.Printf("Data: %s\n", block.Data)
	}
}
