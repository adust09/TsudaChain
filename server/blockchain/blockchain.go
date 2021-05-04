package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
)

// ブロックの構造体
type Block struct {
	Hash          string
	PrevBlockHash string
	Data          string
}

// ブロックチェーンの構造体
type Blockchain struct {
	Blocks []*Block
}

// ブロックの構造体のHashにセットする
func (b *Block) setHash() {
	hash := sha256.Sum256([]byte(b.PrevBlockHash + b.Data))
	b.Hash = hex.EncodeToString(hash[:])
}

// 新しいブロックを作る
func NewBlock(data string, prevBlockHash string) *Block {
	block := &Block{
		Data:          data,
		PrevBlockHash: prevBlockHash,
	}
	block.setHash()

	return block
}

// ブロックチェーンにブロックを追加する
func (bc *Blockchain) AddBlock(data string) *Block {
	prevBlock := bc.Blocks[len(bc.Blocks)-1]
	newBlock := NewBlock(data, prevBlock.Hash)
	bc.Blocks = append(bc.Blocks, newBlock)

	return newBlock
}

// ブロックチェーンを作る
func NewBlockchain() *Blockchain {
	return &Blockchain{[]*Block{NewGenesisBlock()}}
}

func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", "")
}
