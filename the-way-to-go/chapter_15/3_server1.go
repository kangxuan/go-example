package main

import (
	"fmt"
	"net"
	"os"
	"strings"
)

var mapUsers map[string]int

func main() {
	// 服务器
	var (
		listener net.Listener
		error1   error
		conn     net.Conn
	)

	mapUsers = make(map[string]int)

	fmt.Println("开启服务...")

	// 创建 listener
	listener, error1 = net.Listen("tcp", "localhost:50000")
	checkError(error1)

	// 监听来自客户端的请求
	for {
		conn, error1 = listener.Accept()
		checkError(error1)

		go doServerStuff1(conn)
	}
}

func doServerStuff1(conn net.Conn) {
	var buf []byte
	var error1 error

	for true {
		// 用来接收请求的数据
		buf = make([]byte, 512)
		_, error1 = conn.Read(buf)
		checkError2(error1)

		input := string(buf)
		// 关闭服务器
		if strings.Contains(input, ": SH") {
			fmt.Println("服务器已停止...")
			os.Exit(0)
		}

		// 列出已连接服务器的客户端名称列表
		if strings.Contains(input, ": WHO") {
			DisplayList()
		}

		// 截取says之前的数据就是客户端的名称
		ix := strings.Index(input, "says")
		clName := input[0 : ix-1]

		mapUsers[string(clName)] = 1
		fmt.Printf("Received data: --%v--", string(buf))
	}
}

// checkError2 检查错误
func checkError2(err error) {
	if err != nil {
		panic("错误：" + err.Error())
	}
}

func DisplayList() {
	fmt.Println("--------------------------------------------")
	fmt.Println("This is the client list: 1=active, 0=inactive")
	for key, value := range mapUsers {
		fmt.Printf("User %s is %d\n", key, value)
	}
	fmt.Println("--------------------------------------------")
}
