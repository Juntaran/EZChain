/** 
  * Author: Juntaran 
  * Email:  Jacinthmail@gmail.com 
  * Date:   2018/4/22 16:35
  */

package g

import (
	"crypto/sha256"
	"encoding/base64"
)

// 验证某个交易不需要现在整个 block ，而仅仅需要交易 hash 值、Merkle 根节点 hash 值以及 Merkle 路径即可
type MerkleTree struct {
	RootNode 	*MerkleNode
}

type MerkleNode struct {
	Left  		*MerkleNode		// 左孩子
	Right 		*MerkleNode		// 右孩子
	Data  		[]byte			// 内容
	DataShow 	string			// 内容哈希
}


// 创建 Merkle 树
func NewMerkleTree(data [][]byte) *MerkleTree {
	//for _, v := range data {
	//	fmt.Println("NewMerkleTreeData:", string(v))
	//}
	if len(data) == 0 {
		data = append(data, []byte(""))
	}
	var nodes []MerkleNode

	// 叶子节点的数量一定是偶数，然后并非每个 block 都恰好有偶数个交易
	// 当 block 有奇数个交易时，最后一个交易会被复制一次
	if len(data) % 2 != 0 {
		data = append(data, data[len(data)-1])
	}
	for _, v := range data {
		node := NewMerkleNode(nil, nil, v)
		nodes = append(nodes, *node)
	}
	for i := 0; i < len(data)/2; i++ {
		var newLevel []MerkleNode
		for j := 0; j < len(nodes); j += 2 {
			node := NewMerkleNode(&nodes[j], &nodes[j+1], nil)
			newLevel = append(newLevel, *node)
		}
		nodes = newLevel
	}
	mTree := MerkleTree{&nodes[0]}
	return &mTree
}

// 创建新的 Merkle 节点
func NewMerkleNode(left, right *MerkleNode, data []byte) *MerkleNode {
	mNode := MerkleNode{}
	if left == nil && right == nil {
		hash := sha256.Sum256(data)
		mNode.Data = hash[:]
	} else {
		prevHashes := append(left.Data, right.Data...)
		hash := sha256.Sum256(prevHashes)
		mNode.Data = hash[:]
		mNode.DataShow = base64.StdEncoding.EncodeToString(mNode.Data)
	}
	mNode.Left = left
	mNode.Right = right
	return &mNode
}

// Input merkleTree 的根，判断 MerkleTree 的正确性
func JudgeMerkleTree(root *MerkleNode) bool {
	var nodes []*MerkleNode
	// 所有节点入队，依次判断各个节点正确性
	nodes = append(nodes, root)
	var current = 0
	var last = 1
	for current < len(nodes) {
		last = len(nodes)
		for current < last {
			if nodes[current].Left != nil {
				nodes = append(nodes, nodes[current].Left)
			}
			if nodes[current].Right != nil {
				nodes = append(nodes, nodes[current].Right)
			}
			current ++
		}
	}
	for _, v := range nodes {
		if judgeMerkleNode(v) == false {
			return false
		}
	}
	return true
}

// 判断 merkleNode 的正确性
func judgeMerkleNode(node *MerkleNode) bool {
	if node.Left == nil && node.Right == nil {
		return true
	}
	prevHashes := append(node.Left.Data, node.Right.Data...)
	hash := sha256.Sum256(prevHashes)
	if len(hash) != len(node.Data) {
		return false
	}
	for k, v := range hash {
		if v != node.Data[k] {
			return false
		}
	}
	return true
}