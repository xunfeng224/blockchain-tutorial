package main

type BlockChain struct {
	blocks []*Block
}

func NewBlockChain() *BlockChain {
	firstBlock := NewGenesisBlock()
	chain := BlockChain{[]*Block{firstBlock}}
	return &chain
}

func (blockChain *BlockChain) AddBlock(data string) {
	preBlock := blockChain.blocks[len(blockChain.blocks)-1]
	newBlock := NewBlock(data, preBlock.Hash)
	blockChain.blocks = append(blockChain.blocks, newBlock)

}
