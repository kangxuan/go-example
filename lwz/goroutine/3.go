package main

import "fmt"

func main() {
	// 创建一个无缓冲的通道
	ch := make(chan int)
	// 开启一个goroutine函数将ch作为参数传进去
	go func(ch chan int) {
		// 接收通道的数据，未接收到之前一直等待
		rec := <-ch
		fmt.Println(rec)
	}(ch)

	// 发送一个数据，没有接收方时会一直等待
	ch <- 10
}
