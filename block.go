package main

import (
	"fmt"
	"time"
)

type Block struct {
	Timestamp     int64 // Unix time of creation
	Data          []byte
	PrevBlockHash []byte
	Hash          []byte
	Nonce         int
}

func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), []byte(data), prevBlockHash, []byte{}, 0}

	pow := NewProofOfWork(block)
	nonce, hash := pow.Mine()

	block.Hash = hash[:]
	block.Nonce = nonce

	return block
}

func NewGenesisBlock() *Block {
	return NewBlock("Block #0", []byte{})
}

func (b *Block) ToString() string {
	return fmt.Sprintf("{Timestamp:%v Data:%s PrevBlockHash:%x Hash:%x Nonce:%v}",
		b.Timestamp, b.Data, b.PrevBlockHash, b.Hash, b.Nonce)
}
