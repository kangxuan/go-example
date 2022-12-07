package main

import (
	"fmt"
	"time"
)

// 一直往ch通道写入数据
func pump4() chan int {
	ch := make(chan int)
	go func() {
		for i := 0; ; i++ {
			ch <- i
		}
	}()
	return ch
}

// 一直从ch中读取数据
func suck4(ch chan int) {
	for true {
		fmt.Println(<-ch)
	}
}

func main() {
	// 将返回值
	stream := pump4()
	go suck4(stream)
	time.Sleep(time.Second)
}
