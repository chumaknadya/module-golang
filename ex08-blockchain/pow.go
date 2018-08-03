package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math"
	"math/big"
	"strconv"
)

type ProofOfWork struct {
	block      *Block
	target     *big.Int
	targetBits int
}

func NewProofOfWork(b *Block, targetBits int) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-(targetBits*4)))
	pow := &ProofOfWork{b, target, (targetBits * 4)}
	return pow
}

func (pow *ProofOfWork) prepareData(nonce int) []byte {
	data := bytes.Join([][]byte{
		pow.block.PrevBlockHash,
		pow.block.Data,
		IntToHex(pow.block.Timestamp),
		IntToHex(int64(pow.targetBits)),
		IntToHex(int64(nonce)),
	},
		[]byte{},
	)
	return data
}

func (pow *ProofOfWork) Run() []byte {
	var hashInt big.Int
	var hash [32]byte
	nonce := 0

	fmt.Printf("Mining the block containing \"%s\"\n", pow.block.Data)
	for nonce < math.MaxInt64 {
		data := pow.prepareData(nonce)
		hash = sha256.Sum256(data)
		fmt.Printf("\r%x", hash)
		hashInt.SetBytes(hash[:])

		if hashInt.Cmp(pow.target) == -1 {
			break
		} else {
			nonce++
		}
	}
	fmt.Print("\n\n")

	return hash[:]
}

func IntToHex(n int64) []byte {
	return []byte(strconv.FormatInt(n, 16))
}
