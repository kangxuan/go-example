package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 10; i++ {
		//i := i
		go func() {
			fmt.Println(i)
		}()
	}
	time.Sleep(1 * time.Second)
}

//10
//10
//10
//10
//7
//10
//10
//10
//10
//10
// 因为for循环直接开启了goroutine，开启并行，这些goroutine执行有快有慢，而用的i又是公共的i，所以当goroutine执行的时候i是什么就是输出什么
