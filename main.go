package main

import (
	"time"
	"net"
	"fmt"
	"io"
	"bufio"
	"encoding/json"

	"github.com/astaxie/beego"

	"EZChain/blockchain"
	"EZChain/g"
	"EZChain/utils"
	_ "EZChain/routers"
)


func handleConn(conn net.Conn)  {
	//defer conn.Close()
	io.WriteString(conn, "Input Message:")
	scanner := bufio.NewScanner(conn)

	go func() {
		for scanner.Scan() {
			newBlock := blockchain.BC.GenerateBlock(blockchain.BC.Blockchain[len(blockchain.BC.Blockchain)-1], scanner.Text())
			utils.PrintBlock(newBlock)
			io.WriteString(conn, "\nInput Message:")
		}
	}()

	// 模拟接收广播
	go func() {
		for {
			// 每 10 秒接收一次
			time.Sleep(time.Second * 30)
			g.Mutex.Lock()
			output, err := json.Marshal(blockchain.BC.Blockchain)
			if err != nil {
				fmt.Println(err)
			}
			g.Mutex.Unlock()
			io.WriteString(conn, string(output) + "\n\n")
		}
	}()
}

func main() {

	oldBlock := blockchain.Block{
		Index: -1,
		PrevHash: "",
		Timestamp: 0,
		Nonce: 0,
		Hash: "",
		Difficulty: g.Difficulty,
		Info: "old",
		MerkleRootHash: "  ",
	}

	genesisBlock := blockchain.BC.GenerateBlock(oldBlock, "First Block")
	utils.PrintBlock(genesisBlock)

	// 维护链
	go blockchain.BC.TimingFixBlockchain(5)

	// 监听 7190 端口
	go func() {
		listener, err := net.Listen("tcp", ":7190")
		if err != nil {
			fmt.Println(err)
		}
		defer listener.Close()

		for {
			conn, err := listener.Accept()
			if err != nil {
				fmt.Println(err)
			}
			go handleConn(conn)
		}
	}()

	beego.Run()
}