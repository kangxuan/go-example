package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	go func() {
		time.Sleep(4 * time.Second)
		v := <-ch
		fmt.Printf("Received : %d\n", v)
	}()

	fmt.Printf("Sending : 10\n")
	ch <- 10 // 因为是一个无缓存通道，阻塞等待接受者准备好
	fmt.Printf("Sent : 10\n")
}

//Sending : 10
//Sent : 10
//Received : 10
