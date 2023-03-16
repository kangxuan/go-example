package main

import (
	"fmt"
)

func f1(in chan int) {
	fmt.Println(<-in)
}

func main() {
	// 一个没有缓冲的通道，即同步通道
	out := make(chan int)
	go f1(out) // 必须在发送通道前，先开启一个goroutine接收
	//go func() {
	//	out<-2
	//}()
	out <- 2 // 因为程序按顺序执行，out通道只有发送操作，同时没有接受者，则阻塞形成死锁。解决办法如下注释。fatal error: all goroutines are asleep - deadlock!
}
