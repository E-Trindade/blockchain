package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

type Blockchain struct {
	Blocks []*Block
}

type Block struct {
	PreviousHash []byte
	Hash         []byte
	Content      string
	Timestamp    int64
}

func (b *Block) generateHash() {
	block_bytes := bytes.Join([][]byte{
		b.PreviousHash,
		[]byte(strconv.FormatInt(b.Timestamp, 10)),
		[]byte(b.Content),
	}, []byte{})
	h := sha256.Sum256(block_bytes)
	b.Hash = h[:]
}

func NewBlock(content string, prevBlockHash []byte) *Block {
	block := new(Block)
	block.Content = content
	block.PreviousHash = prevBlockHash
	block.Timestamp = time.Now().Unix()

	block.generateHash()
	return block
}

func MakeBlockchain() *Blockchain {
	chain := new(Blockchain)
	genesis := NewBlock("Genesis Block", []byte{})
	chain.Blocks = []*Block{genesis}
	return chain
}

func main() {
	blockchain := MakeBlockchain()

	for i := 0; i < 5; i++ {
		blockchain.Blocks = append(blockchain.Blocks, NewBlock("abc", []byte{123, 3}))
	}

	s, _ := json.Marshal(*blockchain)
	fmt.Println(string(s))

}
