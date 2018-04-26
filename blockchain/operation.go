/** 
  * Author: Juntaran 
  * Email:  Jacinthmail@gmail.com 
  * Date:   2018/3/16 00:05
  */

package blockchain

import (
	"time"
	"fmt"
	"strconv"
	"crypto/sha256"
	"encoding/hex"
	"strings"

	"EZChain/g"
)

// 计算哈希
func CalBlockHash(block Block) string {
	record := strconv.Itoa(int(block.Index)) + strconv.Itoa(int(block.Timestamp)) + strconv.Itoa(int(block.Nonce)) + block.PrevHash
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

// 判断 hash 前缀 0 的个数
func isHashValid(hash string, difficulty int64) bool {
	preZero := strings.Repeat("0", int(difficulty))
	return strings.HasPrefix(hash, preZero)
}

func CalMerkleRoot(bc BChain) string {
	var infos [][]byte
	for _, v := range bc.Blockchain {
		infos = append(infos, []byte(v.Info))
	}
	mr := g.NewMerkleTree(infos)
	return mr.RootNode.DataShow
}

func (blockchain *BChain) GenerateBlock(oldBlock Block, message string) Block {
	var newBlock Block
	var temp Block
	if len(BC.Blockchain) == 0 {
		temp = oldBlock
	} else {
		temp = BC.Blockchain[len(BC.Blockchain) - 1]
	}
	if oldBlock.Index != temp.Index {
		fmt.Println("Resolve Conflict")
	}
	newBlock.Index = temp.Index + 1
 	newBlock.Timestamp = time.Now().Unix()
	newBlock.PrevHash = temp.Hash
	newBlock.Difficulty = g.Difficulty
	newBlock.Info = message
	newBlock.MerkleRootHash = CalMerkleRoot(BC)

	for i := 0; ; i++ {
		newBlock.Nonce = int64(i)
		if !isHashValid(CalBlockHash(newBlock), newBlock.Difficulty) {
			fmt.Printf("\rNonce:%v Hash:%s", newBlock.Nonce, CalBlockHash(newBlock))
			continue
		} else {
			newBlock.Hash = CalBlockHash(newBlock)
			fmt.Printf("\nNonce:%v Hash:%s\n", newBlock.Nonce, CalBlockHash(newBlock))
			break
		}
	}
	BC.Blockchain = append(BC.Blockchain, newBlock)
	return newBlock
}

// 维护链
func fixBlockchain() {
	if len(BC.Blockchain) <= 1 {
		return
	}
	var afterMerge []Block
	for i := 0; i < len(BC.Blockchain) - 1; i++ {
		if BC.Blockchain[i].Index >= BC.Blockchain[i+1].Index {
			afterMerge = append(afterMerge, BC.Blockchain[i])
			BC.Blockchain = append(BC.Blockchain[:i], BC.Blockchain[i+1:]...)
		}
	}

	for _, v := range afterMerge {
		// 重新添加 afterMerge 内的区块
		BC.GenerateBlock(BC.Blockchain[len(BC.Blockchain) - 1], v.Info)
	}
}

// 每隔 duration 维护一次链
func (blockchain *BChain) TimingFixBlockchain(duration time.Duration) {
	ticker := time.NewTicker(time.Second * time.Duration(duration))
	for _ = range ticker.C {
		// 维护链
		fixBlockchain()
	}
	ch := make(chan int)
	<- ch
}