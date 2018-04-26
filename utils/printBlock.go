/** 
  * Author: Juntaran 
  * Email:  Jacinthmail@gmail.com 
  * Date:   2018/4/24 09:48
  */

package utils

import (
	"fmt"

	"EZChain/blockchain"
)

func PrintBlock(block blockchain.Block)  {
	fmt.Println("Block {")
	fmt.Printf("    Index:        %v,\n", block.Index)
	fmt.Printf("    Timestamp:    %v,\n", block.Timestamp)
	fmt.Printf("    Nonce:        %v,\n", block.Nonce)
	fmt.Printf("    MerkleRoot :  \"%s\"\n", block.MerkleRootHash)
	fmt.Printf("    PrevHash:     \"%s\",\n", block.PrevHash)
	fmt.Printf("    Hash:         \"%s\",\n", block.Hash)
	fmt.Printf("    Difficulty:   %v,\n", block.Difficulty)
	fmt.Printf("    Info:         \"%s\"\n", block.Info)

	fmt.Println("}")
}
