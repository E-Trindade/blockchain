package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"math"
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
	Nonce        int64
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

func GenerateProofOfWork(b Block) int64 {
	nonce := int64(1)
	for {
		b.Nonce = nonce
		if b.isValid() {
			return nonce
		}
		nonce += 1
	}
}

func (b *Block) isValid() bool {
	block_bytes := bytes.Join([][]byte{
		b.Hash,
		[]byte(strconv.FormatInt(b.Nonce, 10)),
	}, []byte{})
	h := sha256.Sum256(block_bytes)
	var sum int64 = 0
	for i := 0; i < len(h); i++ {
		sum += int64(i)
	}
	fmt.Println(len(h))
	fmt.Println(sum)
	if sum < math.MaxInt64>>50 {
		return true
	}
	return false
}

func NewBlock(content string, prevBlockHash []byte) *Block {
	block := new(Block)
	block.Content = content
	block.PreviousHash = prevBlockHash
	block.Timestamp = time.Now().Unix()

	block.generateHash()

	block.Nonce = GenerateProofOfWork(*block)

	return block
}

func MakeBlockchain() *Blockchain {
	chain := new(Blockchain)
	genesis := NewBlock("Genesis Block", []byte{})
	chain.Blocks = []*Block{genesis}
	return chain
}

func (bc *Blockchain) mineBlock() {
	bc.Blocks = append(bc.Blocks, NewBlock("abc", bc.Blocks[len(bc.Blocks)-1].Hash))
}

func main() {
	blockchain := MakeBlockchain()

	for i := 0; i < 5; i++ {
		blockchain.mineBlock()
		a, _ := json.Marshal(*(blockchain.Blocks[len(blockchain.Blocks)-1]))
		fmt.Println(string(a))
	}

	s, _ := json.Marshal(*blockchain)
	fmt.Println(string(s))
}
