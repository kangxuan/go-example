package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
	"test_15/rpc_objects"
	"time"
)

func main() {
	calc := new(rpc_objects.Args)
	// 注册该对象
	err := rpc.Register(calc)
	if err != nil {
		return
	}
	// 调用
	rpc.HandleHTTP()

	listener, e := net.Listen("tcp", "localhost:1234")
	if e != nil {
		log.Fatal("Starting RPC-server -listen error:", e)
	}

	// 协程启动，每一个进入listener的HTTP连接创建新的服务线程
	go func() {
		err := http.Serve(listener, nil)
		if err != nil {
			return
		}
	}()
	time.Sleep(1000e9)
}
