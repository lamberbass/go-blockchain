package main

import "fmt"

func main() {
	bc := NewBlockchain()

	bc.AddBlock("Block #1")
	bc.AddBlock("Block #2")

	for _, block := range bc.blocks {
		fmt.Printf("Block: %s ", block.ToString())

		pow := NewProofOfWork(block)
		fmt.Printf("IsValid: %t\n", pow.Validate())
	}
}
