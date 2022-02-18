package main

import (
	"fmt"

	"github.com/6233/jhcoin/blockchain"
)

func main() {
	chain := blockchain.GetBlockchain()
	chain.AddBlock("Second Block")
	chain.AddBlock("Third Block")
	chain.AddBlock("Fourth Block")
	for _, block := range chain.AllBlock() {
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hata: %s\n", block.Hash)
		fmt.Printf("PrevHash: %s\n", block.PrevHash)
	}
}