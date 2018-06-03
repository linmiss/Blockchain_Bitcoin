/*
* this proof of work block chain
 */
package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math"
	"math/big"
)

// Difficulty value, means that the first 24 bits of the hash must be 0
const targetBits = 24

const maxNonce = math.MaxInt64

/*
* The amount of work per Block has to be proven, so there's a pointer to the Block
* we must find the hash < target
 */
type ProofOfWork struct {
	block  *Block
	target *big.Int
}

// target === 1 left move 256-targetBits bit
func NewProofOfWork(b *Block) *ProofOfWork {
	target := big.NewInt(1)
	// Lsh sets z = x << n and returns z.
	target.Lsh(target, uint(256-targetBits))

	pow := &ProofOfWork{b, target}
	return pow
}

//proof of work need params: PreBlockHash, Data, Timestamp, targetBits, nonce
func (pow *ProofOfWork) prepareData(nonce int) []byte {
	data := bytes.Join(
		[][]byte{
			pow.block.PreBlockHash,
			pow.block.HashTransactions(),
			IntToHex(pow.block.Timestamp),
			IntToHex(int64(targetBits)),
			IntToHex(int64(nonce)),
		},
		[]byte{},
	)
	return data
}

// find valid hash
func (pow *ProofOfWork) Run() (int, []byte) {
	var hashInt big.Int
	var hash [32]byte
	nonce := 0

	fmt.Printf("Mining the block containing")
	
	for nonce < maxNonce {
		data := pow.prepareData(nonce)

		hash = sha256.Sum256(data)
		hashInt.SetBytes(hash[:])

		if hashInt.Cmp(pow.target) == -1 {
			fmt.Printf("\r%x", hash)
			break
		} else {
			nonce++
		}
	}
	fmt.Print("\n\n")

	return nonce, hash[:]
}

//valid proof of work : if hash < target hash => valid proof
func (pow *ProofOfWork) Validate() bool {
	var hashInt big.Int

	data := pow.prepareData(pow.block.Nonce)
	hash := sha256.Sum256(data)
	hashInt.SetBytes(hash[:])

	isValid := hashInt.Cmp(pow.target) == -1

	return isValid
}
