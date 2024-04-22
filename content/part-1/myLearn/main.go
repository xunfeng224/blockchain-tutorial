package main

import "fmt"

func main() {
	bc := NewBlockChain()
	bc.AddBlock("This is 1 BTC")
	bc.AddBlock("This is 2 BTC")
	for i, block := range bc.blocks {
		fmt.Printf("第%d个Block:\n", i)
		fmt.Printf("Data:%s\n", block.Data)
		fmt.Printf("Prev hash:%x\n", block.PrevBlockHash)
		fmt.Printf("Hash:%x\n", block.Hash)
		fmt.Println()
	}
}
