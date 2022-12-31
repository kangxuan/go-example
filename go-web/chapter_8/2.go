package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
)

// 展示TCP client搭建
func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s host:port ", os.Args[0])
		os.Exit(1)
	}
	service := os.Args[1]
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)
	// 连接服务
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkError(err)
	_, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	checkError(err)
	result, err := ioutil.ReadAll(conn)
	checkError(err)
	fmt.Println(string(result))
	os.Exit(0)
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

//go run 2.go baidu.com:80
//HTTP/1.1 200 OK
//Date: Sat, 31 Dec 2022 01:46:39 GMT
//Server: Apache
//Last-Modified: Tue, 12 Jan 2010 13:48:00 GMT
//ETag: "51-47cf7e6ee8400"
//Accept-Ranges: bytes
//Content-Length: 81
//Cache-Control: max-age=86400
//Expires: Sun, 01 Jan 2023 01:46:39 GMT
//Connection: Close
//Content-Type: text/html
