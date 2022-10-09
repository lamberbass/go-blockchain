package main

import (
	"math/big"
)

func Int64ToBytes(number int64) []byte {
	big := new(big.Int)
	return big.SetInt64(number).Bytes()
}
