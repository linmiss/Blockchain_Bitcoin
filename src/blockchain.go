package main

// BlockChain is a pointer array of Block
type BlockChain struct {
	blocks []*Block
}

//create a genesis chain
func NewBlockChain() *BlockChain {
	return &BlockChain{[]*Block{NewGenesisBlock()}}
}

//add new block to chain
//data: trading
func (bc *BlockChain) AddBlock(data string) {
	preBlock := bc.blocks[len(bc.blocks)-1]
	newBlock := NewBlock(data, preBlock.Hash)
	bc.blocks = append(bc.blocks, newBlock)
}
