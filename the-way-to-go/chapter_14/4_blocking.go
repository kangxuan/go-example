package main

import (
	"fmt"
)

func f1(in chan int) {
	fmt.Println(<-in)
}

func main() {
	out := make(chan int)
	out <- 2 // 因为程序按顺序执行，out通道只有发送操作，同时没有接受者，则阻塞形成死锁。解决办法如下注释。
	//go func() {
	//	out<-2
	//}()
	go f1(out)
}
