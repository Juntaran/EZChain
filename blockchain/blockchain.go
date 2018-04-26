/** 
  * Author: Juntaran 
  * Email:  Jacinthmail@gmail.com 
  * Date:   2018/3/15 16:22
  */

package blockchain

type Transaction struct {
	Sender				string			// 发送方
	Reciever			string			// 接收方
	Amount				int64			// 交易量
}

type Block struct {
	Index				int64			// 当前 block 在 blockchain 中的位置
	PrevHash 			string			// 前一个 block 的 hash
	Timestamp 			int64			// blcok 生成的时间戳
	Nonce 				int64			// 随机数
	Hash 				string			// 当前 block 的 hash
	Difficulty			int64			// 当前区块链生成难度
	Info 				string			// 传递的值
	//Transaction			Transaction		// 记录这个区块发生的交易
	MerkleRootHash		string			// Merkle 树根结点
}

type BChain struct {
	Blockchain 		[]Block				// block 串成 chain
}

var BC BChain