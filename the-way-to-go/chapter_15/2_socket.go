package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	var (
		host   = "www.apache.org"
		port   = "80"
		remote = host + ":" + port
		msg    = "GET / \n"
		data   = make([]uint8, 4096)
		read   = true
		count  = 0
	)

	// 创建一个socket
	con, err := net.Dial("tcp", remote)
	// 检查连接器
	checkConnection(con, err)
	// 发送消息
	_, err = io.WriteString(con, msg)
	if err != nil {
		return
	}

	for read {
		count, err = con.Read(data)
		read = err == nil
		fmt.Println(string(data[0:count]))
	}
	// 关闭socket
	err1 := con.Close()
	if err1 != nil {
		return
	}
}

// checkConnection 检查链接
func checkConnection(conn net.Conn, err error) {
	if err != nil {
		fmt.Printf("error %v connecting!", err)
		os.Exit(1)
	}
	fmt.Printf("Connection is made with %v\n", conn)
}

//Connection is made with &{{0xc00011a000}}
