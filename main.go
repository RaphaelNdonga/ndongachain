package main

import (
	"fmt"
	"strconv"

	"github.com/RaphaelNdonga/ndongachain/blockchain"
)

func main() {
	chain := blockchain.InitBlockChain()

	chain.AddBlock("First Block after genesis")
	chain.AddBlock("Second Block after genesis")
	chain.AddBlock("Third Block after genesis")

	for _, block := range chain.Blocks {
		fmt.Printf("Current Hash: %x\n", block.Hash)
		fmt.Printf("Prev Hash: %x\n ", block.PrevHash)
		fmt.Printf("Data: %s\n", block.Data)

		pow := blockchain.NewProof(block)
		fmt.Printf("pow: %s", strconv.FormatBool(pow.Validate()))
		fmt.Println()
	}
}
