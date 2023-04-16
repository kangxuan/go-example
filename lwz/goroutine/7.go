package main

import "fmt"

// Producer2 返回一个只接收通道，单向通道一般用于保障数据安全
func Producer2() <-chan int {
	ch := make(chan int, 10)
	go func() {
		for i := 0; i < 10; i++ {
			if i%2 == 1 {
				ch <- i
			}
		}

		close(ch)
	}()

	return ch
}

// Consumer2 参数为只接收通道
func Consumer2(ch <-chan int) int {
	Sum := 0
	for v := range ch {
		Sum += v
	}

	return Sum
}

func main() {
	ch2 := Producer2()
	res2 := Consumer2(ch2)
	fmt.Println(res2)
}
