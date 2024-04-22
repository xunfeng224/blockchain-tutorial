package main

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

type Block struct {
	Timestamp     int64
	Data          []byte
	PrevBlockHash []byte
	Hash          []byte
}

func NewBlock(data string, prevBlockHash []byte) *Block {
	block := Block{
		time.Now().Unix(),
		[]byte(data),
		prevBlockHash,
		[]byte{}}
	block.SetHash()
	return &block
}

func (block *Block) SetHash() {
	timeStamp := []byte(strconv.FormatInt(block.Timestamp, 10))
	header := bytes.Join([][]byte{timeStamp, block.Data, block.PrevBlockHash}, []byte{})
	hash := sha256.Sum256(header)
	block.Hash = hash[:]
}

func NewGenesisBlock() *Block{
	genesisBlock := NewBlock("First Block", []byte{})
	return genesisBlock
}
