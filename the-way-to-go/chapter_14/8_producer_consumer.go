package main

import (
	"fmt"
)

func S(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i * 10
	}
}

func G(ch chan int, done chan bool) {
	for true {
		fmt.Printf("%d\n", <-ch)
	}
	done <- true
}

func main() {
	// 有两个协程，第一个提供数字 0，10，20，...90 并将他们放入通道，第二个协程从通道中读取并打印
	ch := make(chan int)
	done := make(chan bool)
	go S(ch)
	go G(ch, done)
	<-done
}
