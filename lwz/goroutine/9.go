package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan bool)
	select {
	case v := <-c:
		fmt.Println(v)
	case <-time.After(3 * time.Second): // 等待3秒
		fmt.Println("timeout")
	}
}
