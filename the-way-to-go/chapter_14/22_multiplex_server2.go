package main

import (
	"fmt"
)

type Request3 struct {
	a, b   int
	replyc chan int // 定义一个通道
}

func (r *Request3) String() string {
	return fmt.Sprintf("%d + %d = %d", r.a, r.b, <-r.replyc)
}

type binOp func(a, b int) int

func run(op binOp, req *Request3) {
	req.replyc <- op(req.a, req.b)
}

func server3(op binOp, service chan *Request3, quit chan bool) {
	// 无限循环监听Request是否有发送数据
	for {
		select {
		// 从service中读取数据
		case req := <-service:
			go run(op, req)
		// 当接收到quit则退出
		case <-quit:
			return
		}
	}
}

func startServer(op binOp) (chan *Request3, chan bool) {
	reqChan := make(chan *Request3)
	quit := make(chan bool)
	go server3(op, reqChan, quit)
	return reqChan, quit
}

func main() {
	adder, quit := startServer(func(a, b int) int {
		return a + b
	})

	req1 := &Request3{3, 4, make(chan int)}
	req2 := &Request3{150, 250, make(chan int)}

	adder <- req1
	adder <- req2

	fmt.Println(req1, req2)

	quit <- true
	fmt.Println("done")
}

//3 + 4 = 7 150 + 250 = 400
//done
