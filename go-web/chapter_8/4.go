package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

// 展示用goroutine优化TCP server
func main() {
	service := ":8888"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError2(err)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError2(err)
	for {
		conn, err := listener.Accept()
		if err != nil {
			// 一般还有日志记录
			continue
		}
		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()
	daytime := time.Now().String()
	conn.Write([]byte(daytime))
}

func checkError2(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
