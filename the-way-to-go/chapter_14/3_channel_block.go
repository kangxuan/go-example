package main

import (
	"fmt"
	"time"
)

func main() {
	//runtime.GOMAXPROCS(4)
	ch1 := make(chan int)
	go pump3(ch1)
	go suck3(ch1)
	time.Sleep(1 * time.Second)
}

// 通道提供数值，也叫生产者
func pump3(ch chan int) {
	for i := 0; ; i++ {
		ch <- i
	}
}

// 读取通道数据，也叫消费者
func suck3(ch chan int) {
	for true {
		fmt.Println(<-ch)
	}
}
