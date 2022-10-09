package main

import "fmt"

func main() {
	bc := NewBlockchain()

	bc.AddBlock("Block #1")
	bc.AddBlock("Block #2")

	for _, block := range bc.blocks {
		fmt.Println(block.ToString())
	}
}
