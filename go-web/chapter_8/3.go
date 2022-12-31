package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

// 展示TCP server端搭建
func main() {
	service := ":7777"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError1(err)
	// 服务监听
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError1(err)
	// 接收请求
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		daytime := time.Now().String()
		conn.Write([]byte(daytime))
		conn.Close()
	}
}

func checkError1(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
