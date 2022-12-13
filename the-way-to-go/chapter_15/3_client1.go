package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	// 客户端
	var conn net.Conn
	var error1 error
	var inputReader *bufio.Reader
	var input string
	var clientName string

	conn, error1 = net.Dial("tcp", "localhost:50000")
	checkError1(error1)

	inputReader = bufio.NewReader(os.Stdin)
	fmt.Println("先输入您的姓名")
	clientName, _ = inputReader.ReadString('\n')
	trimClient := strings.Trim(clientName, "\n")

	for true {
		fmt.Printf("请发送消息给服务器? Q - 断开连接. SH - 关闭服务器. SHO")
		input, _ = inputReader.ReadString('\n')
		trimInput := strings.Trim(input, "\n")

		if trimInput == "Q" {
			fmt.Println("已断开连接...")
			return
		}

		_, error1 = conn.Write([]byte(trimClient + " says: " + trimInput))
	}
}

func checkError1(error error) {
	if error != nil {
		panic("Error: " + error.Error()) // terminate program
	}
}
