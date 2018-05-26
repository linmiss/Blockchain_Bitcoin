package main

import (
	"fmt"
)

func main() {
	bc := NewBlockChain()

	bc.AddBlock("send first BTC to May")
	bc.AddBlock("send second BTC to May")

	for index, block := range bc.blocks {
		fmt.Println("----------------------", index, "block")
		fmt.Printf("PreBlockHash: %x\n", block.PreBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
	}
}
