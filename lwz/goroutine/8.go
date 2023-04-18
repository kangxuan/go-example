package main

import "fmt"

func main() {
	ch := make(chan int, 2)
	for i := 0; i < 10; i++ {
		select {
		case w := <-ch: //当ch有数据时取出
			fmt.Println(w)
		case ch <- i: //当ch无数据时写入数据
		}
	}
}
