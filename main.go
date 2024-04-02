package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
)

type Block struct {
	Index     int
	Timestamp string
	Data      string
	PrevHash  string
	Hash      string
	Nonce     int
}

// 查询区块链
func getBlockchain() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error connecting:", err.Error())
		return
	}
	defer conn.Close()

	// 发送消息到服务器
	conn.Write([]byte("GetBlockchain"))

	// 接收区块链数据
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
		return
	}
	var chain []Block
	json.Unmarshal(buf[:n], &chain)

	// 打印区块链
	chainInBytes, _ := json.MarshalIndent(chain, "", "  ")
	fmt.Printf("Blockchain:\n%s\n", string(chainInBytes))
}

func addBlockToBlockchain(data string) {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error connecting:", err.Error())
		return
	}
	defer conn.Close()

	// 发送添加区块的消息到服务器
	conn.Write([]byte("AddBlock|" + data))

	// 接收服务器响应
	reader := bufio.NewReader(conn)
	response, err := reader.ReadString('\r')
	if err != nil {
		fmt.Println("Error reading response:", err.Error())
		return
	}
	fmt.Println("Server response:", response)
}

// 向服务器发送添加区块的请求，并执行拜占庭容错共识算法
func addBlockToBlockchainWithBFT(data string) {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error connecting:", err.Error())
		return
	}
	defer conn.Close()

	// 发送添加区块的消息到服务器
	conn.Write([]byte("AddBlockWithBFT|" + data))

	// 接收服务器响应
	reader := bufio.NewReader(conn)
	response, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading response:", err.Error())
		return
	}
	fmt.Println("Server response:", response)
}

// 连接服务器并发送加入请求
func joinNetwork(ip string, port int) {
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", ip, port))
	if err != nil {
		fmt.Println("Error connecting:", err.Error())
		return
	}
	defer conn.Close()

	// 发送加入请求，包含自己的节点信息
	conn.Write([]byte(fmt.Sprintf("Join|%s|%d", "127.0.0.1", 8080))) // 举例，假设节点 IP 是 127.0.0.1，端口是 8081

	// 接收服务器响应
	reader := bufio.NewReader(conn)
	response, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading response:", err.Error())
		return
	}
	fmt.Println("Server response:", response)
}

func main() {
	joinNetwork("127.0.0.1", 8080)
	//getBlockchain()

	//reader := bufio.NewReader(os.Stdin)
	//fmt.Print("Enter data for new block: ")
	//data, _ := reader.ReadString('\r')
	//
	//// 去除末尾的换行符
	//data = data[:len(data)-1]
	//
	//// 添加区块到区块链
	//addBlockToBlockchain(data)
}
