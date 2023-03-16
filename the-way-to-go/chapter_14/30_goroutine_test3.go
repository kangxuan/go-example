package main

import "fmt"

func f2(ch chan int) {
	for {
		// 多返回值
		// v是值如果通道被关闭则返回对应类型的零值，
		// ok通道ch关闭时返回 false，否则返回 true
		v, ok := <-ch
		if !ok {
			fmt.Println("ch closed!")
			break
		}

		fmt.Println(v)
	}
}

func f3(ch chan int) {
	// 通过for range更常用获取通道数据
	for v := range ch {
		fmt.Println(v)
	}
}

func main() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2

	// close chan，通道关闭后依然能够取出通道里面的数据
	close(ch)
	//f2(ch)
	//fmt.Print("success")

	f3(ch)
	fmt.Println("success1")
}
