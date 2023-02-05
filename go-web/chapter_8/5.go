package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

// 展示TCP server接收处理client的请求数据
func main() {
	service := ":9999"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError3(err)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError3(err)
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go handleClient3(conn)
	}
}

// 处理客户端
func handleClient3(conn net.Conn) {
	conn.SetReadDeadline(time.Now().Add(2 * time.Minute)) // 设置2分钟的超时时间，当一定时间内客户端无请求发送，conn 便会自动关闭
	request := make([]byte, 128)
	defer conn.Close()
	for {
		readLen, err := conn.Read(request)
		if err != nil {
			fmt.Println(err)
			break
		}
		if readLen == 0 {
			break
		} else if strings.TrimSpace(string(request[:readLen])) == "timestamp" {
			daytime := strconv.FormatInt(time.Now().Unix(), 10)
			conn.Write([]byte(daytime))
		} else {
			daytime := time.Now().String()
			conn.Write([]byte(daytime))
		}

		request = make([]byte, 128)
	}
}

// 检查错误
func checkError3(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
