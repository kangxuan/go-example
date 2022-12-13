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
	// 创建连接器
	conn, err := net.Dial("tcp", "localhost:50000")
	if err != nil {
		fmt.Println("创建连接器失败", err.Error())
		// 终止程序
		return
	}

	// 创建屏幕读取器
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("请输入您的名字？")
	// 读取数据并以\n结束
	clientName, _ := inputReader.ReadString('\n')
	trimmedClient := strings.Trim(clientName, "\n")

	// 向服务器发送数据直到程序退出
	for true {
		fmt.Println("请向服务器发送您的数据？输入Q结束")
		input, _ := inputReader.ReadString('\n')
		trimmedInput := strings.Trim(input, "\n")
		if trimmedInput == "Q" {
			return
		}

		_, err = conn.Write([]byte(trimmedClient + " says: " + trimmedInput))
	}
}

//请输入您的名字？
//shanla
//请向服务器发送您的数据？输入Q结束
//您好！
//请向服务器发送您的数据？输入Q结束
