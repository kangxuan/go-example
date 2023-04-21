package main

import (
	"fmt"
	"net"
)

func main() {
	// 服务端
	fmt.Println("开启服务...")
	// 1. 创建 listener
	listener, err := net.Listen("tcp", "localhost:50000")
	defer listener.Close()
	if err != nil {
		fmt.Println("Error listening", err.Error())
		return
	}
	// 2. 监听并接受来自客户端的连接
	for true {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("错误的请求", err.Error())
		}
		// 3. 给每一个请求单独开启一个goroutine处理
		go doServerStuff(conn)
	}
}

// doServerStuff
func doServerStuff(conn net.Conn) {
	// 关闭连接
	defer conn.Close()
	for true {
		// 定义接受多大的数据
		buf := make([]byte, 512)
		// 读取
		len1, err := conn.Read(buf)
		if err != nil {
			fmt.Println("读取数据错误", err.Error())
			return
		}
		fmt.Printf("接受到的数据：%v\n", string(buf[:len1]))
	}
}

//开启服务...
//接受到的数据：shanla says: 您好！
