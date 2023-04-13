package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 10000; i++ {
		go func() {
			fmt.Println(i)
		}()
	}
	time.Sleep(time.Second)
}