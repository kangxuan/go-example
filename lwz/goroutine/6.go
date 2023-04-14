package main

import (
	"fmt"
)

func main() {
	ch := make(chan int, 5)

	for i := 0; i < 5; i++ {
		ch <- i
	}
	f(ch)
}

func f(ch chan int) {
	for v := range ch {
		fmt.Println(v)
	}
}
