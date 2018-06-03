package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"log"
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
	Transactions []*Transaction
	Nonce        int
}

// make Block to []byte
func (b *Block) Serialize() []byte {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)

	err := encoder.Encode(b)
	if err != nil {
		log.Panic(err)
	}

	return result.Bytes()
}

//reverse Serialize
func DeserializeBlock(d []byte) *Block {
	var block Block
	decoder := gob.NewDecoder(bytes.NewReader(d))

	err := decoder.Decode(&block)
	if err != nil {
		log.Panic(err)
	}

	return &block
}

// NewBlock: product new block (need Data and PreBlockHash as params)
func NewBlock(transactions []*Transaction, PreBlockHash []byte) *Block {
	block := &Block{
		Timestamp:    time.Now().Unix(),
		PreBlockHash: PreBlockHash,
		Hash:         []byte{},
		Transactions: transactions,
		Nonce:        0,
	}

	pow := NewProofOfWork(block)
	nonce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce

	return block
}

// NewGenesisBlock: first block
func NewGenesisBlock(coinbase *Transaction) *Block {
	return NewBlock([]*Transaction{coinbase}, []byte{})
}

// 计算区块里所有交易的哈希
func (b *Block) HashTransactions() []byte {
	var txHashes [][]byte
	var txHash [32]byte

	for _, tx := range b.Transactions {
		txHashes = append(txHashes, tx.ID)
	}
	txHash = sha256.Sum256(bytes.Join(txHashes, []byte{}))

	return txHash[:]
}
