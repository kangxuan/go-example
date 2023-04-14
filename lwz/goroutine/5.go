package main

import "fmt"

func main() {
	ch := make(chan int, 2)
	ch <- 10
	ch <- 20
	close(ch)
	// 关闭通道
	func() {
		for {
			h, ok := <-ch
			if !ok {
				fmt.Printf("通道被关闭，拿到的值：%d\n", h)
				break
			} else {
				fmt.Printf("通道未关闭或还有值，拿到的值：%d\n", h)
			}
		}
	}()
}
