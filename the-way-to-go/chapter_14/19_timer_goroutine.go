package main

import (
	"fmt"
	"time"
)

func main() {
	// 每隔1秒流出一个"当前时间"
	tick := time.Tick(time.Second)
	// 5秒后返回流出一个"当前时间"
	boom := time.After(5 * time.Second)

	for true {
		select {
		case ct := <-tick:
			fmt.Println(ct)
		case ct := <-boom:
			fmt.Println(ct)
			return
		default:
			fmt.Println("    .")
			time.Sleep(time.Second)
		}
	}
}

//    .
//tick.
//    .
//tick.
//    .
//tick.
//    .
//tick.
//    .
//boom.
