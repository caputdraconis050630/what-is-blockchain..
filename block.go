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

func (b *Block) SetHash() {
	timetamp := []byte(strconv.FormatInt(b.Timestamp, 10))                       // 타임스탬프는 기존에 문자열. 이를 10진수로 변환 후, 바이트 배열로 변환
	headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timetamp}, []byte{}) // 이전 블록의 해시, 데이터, 타임스탬프를 합친다.
	hash := sha256.Sum256(headers)                                               // 합친 데이터를 해시한다.
	b.Hash = hash[:]                                                             // 해시값을 블록의 해시값으로 설정한다.
}

func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), []byte(data), prevBlockHash, []byte{}} // 블록을 생성하고, 필요한 데이터를 채운다.
	block.SetHash()
	return block
}
