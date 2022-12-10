package main

import (
	"fmt"
	"time"
)

func main() {
	//runtime.GOMAXPROCS(2)
	ch1 := make(chan int)
	ch2 := make(chan int)

	go pump1(ch1)
	go pump2(ch2)
	go suck1(ch1, ch2)

	time.Sleep(time.Second)
}

// pump1 循环往通道写i+2
func pump1(ch chan int) {
	for i := 0; ; i++ {
		ch <- i + 2
	}
}

// pump2 循环往通道写i+5
func pump2(ch chan int) {
	for i := 0; ; i++ {
		ch <- i + 5
	}
}

// suck1 取数据
func suck1(ch1, ch2 chan int) {
	for i := 0; ; i++ {
		select {
		case v := <-ch1:
			fmt.Printf("%d - Received on channel 1: %d\n", i, v)
		case v := <-ch2:
			fmt.Printf("%d - Received on channel 2: %d\n", i, v)
		}
	}
}
