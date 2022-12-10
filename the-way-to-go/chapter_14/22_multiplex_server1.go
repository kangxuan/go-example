package main

import (
	"fmt"
	"time"
)

type Request2 struct {
	a, b   int
	replyc chan int // 定义一个通道
}

type binOp2 func(a, b int) int

func run2(op binOp2, req *Request2) {
	req.replyc <- op(req.a, req.b)
}

func server2(op binOp2, service chan *Request2, quit chan bool) {
	// 无限循环监听Request是否有发送数据
	for {
		select {
		// 从service中读取数据
		case req := <-service:
			go run2(op, req)
		// 当接收到quit则退出
		case <-quit:
			return
		}
	}
}

func startServer2(op binOp2) (chan *Request2, chan bool) {
	reqChan := make(chan *Request2)
	quit := make(chan bool)
	go server2(op, reqChan, quit)
	return reqChan, quit
}

func main() {
	startTime := time.Now().Second()
	adder, quit := startServer2(func(a, b int) int {
		return a + b
	})
	const N = 1000000
	var reqs [N]Request2
	for i := 0; i < N; i++ {
		req := &reqs[i]
		req.a = i
		req.b = i + N
		req.replyc = make(chan int)
		adder <- req
	}

	for i := N - 1; i >= 0; i-- {
		if <-reqs[i].replyc != N+2*i {
			fmt.Println("fail at", i)
		} else {
			fmt.Println("Request ", i, " is ok!")
		}
	}
	fmt.Println("done")
	endTime := time.Now().Second()
	useTime := endTime - startTime
	quit <- true
	fmt.Println(useTime)
}
