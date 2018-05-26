package main

import (
	"time"
)

// Block: block header and trading
// Timestamp: PreBlockHash, Hash is of block header
//PreBlockHash: previous block hash
//Hash: current block hash
//Data: block information
type Block struct {
	Timestamp    int64
	PreBlockHash []byte
	Hash         []byte
	Data         []byte
	Nonce        int
}

// NewBlock: product new block (need Data and PreBlockHash as params)
func NewBlock(data string, PreBlockHash []byte) *Block {
	block := &Block{
		Timestamp:    time.Now().Unix(),
		PreBlockHash: PreBlockHash,
		Hash:         []byte{},
		Data:         []byte(data),
		Nonce:        0,
	}

	pow := NewProofOfWork(block)
	nonce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce

	return block
}

//SetHash: set current block hash
/*
func (b *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	headers := bytes.Join([][]byte{b.PreBlockHash, b.Data, timestamp}, []byte{})
	hash := sha256.Sum256(headers)
	b.Hash = hash[:]
}
*/

// NewGenesisBlock: first block
func NewGenesisBlock() *Block {
	return NewBlock("Tim's Genesis Block", []byte{})
}
