package main

import "fmt"

func main() {
	bc := NewBlockchain() // 블록체인을 생성한다.

	bc.AddBlock("Send 1 BTC to caput")
	bc.AddBlock("Send 2 more BTC to caput")

	for _, block := range bc.blocks {
		fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Println()
	}
}
