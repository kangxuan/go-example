package main

import (
	"fmt"
	"time"
)

type Request1 struct {
	a, b   int
	replyc chan int // 定义一个通道
}

type binOp1 func(a, b int) int

func run1(op binOp1, req *Request1) {
	req.replyc <- op(req.a, req.b)
}

func server1(op binOp1, service chan *Request1) {
	// 无限循环监听Request是否有发送数据
	for {
		// 从service中读取数据
		req := <-service
		go run1(op, req)
	}
}

func startServer1(op binOp1) chan *Request1 {
	reqChan := make(chan *Request1)
	go server1(op, reqChan)
	return reqChan
}

func main() {
	startTime := time.Now().Second()
	adder := startServer1(func(a, b int) int {
		return a + b
	})
	const N = 1000000
	var reqs [N]Request1
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
	fmt.Println(useTime)
}
