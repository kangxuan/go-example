package main

import (
	"fmt"
	"time"
)

func main() {
	var ch chan int
	for i := 0; i < 3; i++ {
		go func(idx int) {
			ch <- idx
		}(i)
	}

	fmt.Println("result:", <-ch)
	time.Sleep(2 * time.Second)
}
