package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math"
	"math/big"
)

const maxNonce = math.MaxInt64

// Difficulty of mining. For example, when targetBits is 16,
// a valid hash should have at least 4 leading zeros (16 bits)
const targetBits = 16

type ProofOfWork struct {
	block  *Block
	target *big.Int
}

func NewProofOfWork(b *Block) *ProofOfWork {
	target := getTarget()
	pow := &ProofOfWork{b, target}
	return pow
}

func (pow *ProofOfWork) Mine() (int, []byte) {
	var hashInt big.Int
	var hash [32]byte
	nonce := 0

	fmt.Printf("Mining the block containing data: %s\n", pow.block.Data)

	for nonce < maxNonce {
		hash = generateSha256Hash(pow, nonce)
		hashInt.SetBytes(hash[:])

		if hashInt.Cmp(pow.target) == -1 {
			break
		} else {
			nonce++
		}
	}

	return nonce, hash[:]
}

func (pow *ProofOfWork) Validate() bool {
	var hashInt big.Int

	hash := generateSha256Hash(pow, pow.block.Nonce)
	hashInt.SetBytes(hash[:])

	isValid := hashInt.Cmp(pow.target) == -1
	return isValid
}

func generateSha256Hash(pow *ProofOfWork, nonce int) [32]byte {
	bytes := bytes.Join(
		[][]byte{
			pow.block.PrevBlockHash,
			pow.block.Data,
			Int64ToBytes(pow.block.Timestamp),
			pow.target.Bytes(),
			Int64ToBytes(int64(nonce)),
		},
		[]byte{},
	)
	hash := sha256.Sum256(bytes)
	return hash
}

func getTarget() *big.Int {
	// Initialize a big.Int with the value of 1
	target := big.NewInt(1)

	// Shift it left by 256 - targetBits (256 is the length of a SHA-256 hash in bits)
	target.Lsh(target, uint(256-targetBits))

	// Example: when targetBits is 16 (2 bytes), the target is 0001000000000000000000000000000000000000000000000000000000000000
	// A valid hash should be smaller than the target, so it should have at least 4 leading zeros (2 bytes)
	return target
}
