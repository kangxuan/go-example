package main

import "fmt"

func main() {
	// 当给通道设置了容量，只发送信息给通道时，不会有问题，因为有一个位置
	ch := make(chan int, 1)
	ch <- 1
	//ch <- 2 // 再发送给通道就有问题了
	fmt.Println("success")
}
