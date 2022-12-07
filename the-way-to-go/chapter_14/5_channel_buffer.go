package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 1)
	go func() {
		time.Sleep(4 * time.Second)
		v := <-ch
		fmt.Printf("The Received : %d\n", v)
	}()

	fmt.Println("Sending 100")
	ch <- 100 // 因为是带缓存的通道，此时写入不会有阻塞，程序则继续执行不会等待接收者
	fmt.Println("Sent 100")
}
