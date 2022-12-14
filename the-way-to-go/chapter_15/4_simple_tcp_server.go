package main

import (
	"flag"
	"fmt"
	"net"
	"syscall"
)

var maxRead = 25

func main() {
	// 读取命令行参数
	flag.Parse()
	if flag.NArg() != 2 {
		panic("请输入监听地址和端口")
	}

	// 组装hostAndPort
	hostAndPort := fmt.Sprintf("%s:%s", flag.Arg(0), flag.Arg(1))
	// 初始化服务器
	listener := initServer(hostAndPort)
	for true {
		conn, err := listener.Accept()
		checkError3(err, "接收连接: ")
		// 每个链接都会以协程的方式运行
		go connectionHandler(conn)
	}
}

// 初始化服务器
func initServer(hostAndPort string) net.Listener {
	// 得到服务器地址和端口
	serverAddr, err := net.ResolveTCPAddr("tcp", hostAndPort)
	checkError3(err, "Resolving address:port failed: '"+hostAndPort+"'")

	listener, err := net.ListenTCP("tcp", serverAddr)
	checkError3(err, "ListenTCP: ")
	println("Listening to: ", listener.Addr().String())
	return listener
}

func connectionHandler(conn net.Conn) {
	// 返回客户端的地址
	connFrom := conn.RemoteAddr().String()
	println("连接来自：" + connFrom)
	sayHello(conn)
	for {
		// 初始化数据读取buf
		var ibuf = make([]byte, maxRead+1)
		// 读取数据
		length, err := conn.Read(ibuf[0:maxRead])
		ibuf[maxRead] = 0

		// 处理消息
		switch err {
		case nil:
			handleMsg(length, ibuf)
		case syscall.EAGAIN: // try again
			continue
		default:
			goto DISCONNECT
		}
	}

DISCONNECT:
	err := conn.Close()
	println("断开连接：", connFrom)
	checkError3(err, "关闭：")
}

func handleMsg(length int, msg []byte) {
	if length > 0 {
		print("<", length, ":")
		for i := 0; ; i++ {
			if msg[i] == 0 {
				break
			}
			fmt.Printf("%c", msg[i])
		}
		print(">")
	}
}

func sayHello(conn net.Conn) {
	obuf := []byte{'L', 'e', 't', '\'', 's', ' ', 'G', 'O', '!', '\n'}
	// 向客户端广播
	wrote, err := conn.Write(obuf)
	checkError3(err, "Write: wrote "+string(rune(wrote))+" bytes.")
}

// 检查错误
func checkError3(err error, info string) {
	if err != nil {
		panic("错误：" + info + " " + err.Error())
	}
}
