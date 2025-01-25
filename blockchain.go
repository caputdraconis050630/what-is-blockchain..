package main

type Blockchain struct {
	blocks []*Block
}

// AddBlock saves provided data as a block in the blockchain
// AddBlock은 블록체인에 블록을 추가한다.
func (bc *Blockchain) AddBlock(data string) {
	prevBlock := bc.blocks[len(bc.blocks)-1]
	newBlock := NewBlock(data, prevBlock.Hash)
	bc.blocks = append(bc.blocks, newBlock)
}

// NewBlockchain creates a new Blockchain with genesis Block
// NewBlockchain은 체인을 생성하고, 제네시스 블록을 추가한다.
func NewBlockchain() *Blockchain {
	return &Blockchain{[]*Block{NewGenesisBlock()}}
}
