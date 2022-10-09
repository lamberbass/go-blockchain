package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"strconv"
	"time"
)

type Block struct {
	Timestamp     int64 // Unix time of creation
	Data          []byte
	PrevBlockHash []byte
	Hash          []byte
}

// Concatenates the block fields and calculates a SHA-256 hash
func (b *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{})
	hash := sha256.Sum256(headers)

	b.Hash = hash[:]
}

func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), []byte(data), prevBlockHash, []byte{}}
	block.SetHash()
	return block
}

func NewGenesisBlock() *Block {
	return NewBlock("Block #0", []byte{})
}

func (b *Block) ToString() string {
	return fmt.Sprintf("{Timestamp:%v Data:%s PrevBlockHash:%x Hash:%x}", b.Timestamp, b.Data, b.PrevBlockHash, b.Hash)
}
