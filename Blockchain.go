package main

import (
	"crypto/sha256"
	"fmt"
	"time"
)

type Blockchain struct {
	blocks []*Block
}

type Block struct {
	previousHash int64
	hash         int64
	timestamp    time.Time
	content      string
}

func (b *Block) generateHash() {
	hash := sha256.Sum256([]byte{
		byte(b.previousHash),
		byte(b.timestamp.Unix()),
		byte(b.content),
	})
	b.hash = hash
}
func (b *Blockchain) createGenesis() {
	genesis := new(Block)
	genesis.previousHash = 0
	genesis.hash = 123
	genesis.timestamp = time.Now()
	b.blocks = []*Block{genesis}
}

func makeBlockchain() *Blockchain {
	chain := new(Blockchain)
	chain.createGenesis()
	return chain
}

func main() {
	blockchain := makeBlockchain()
	fmt.Printf("%#v\n", *blockchain)
}
