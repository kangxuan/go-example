package main

import "fmt"

func main() {
	// 创建一个容量为1的有缓冲区通道
	ch := make(chan int, 1)
	// 发送时不会阻塞，因为有1个容量
	ch <- 10
	// 再发送一个就会阻塞，因为容量已经满了
	//ch <- 10
	fmt.Println("发送成功")
}
