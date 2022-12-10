package main

import "fmt"

// tel
func tel(ch chan int) {
	for i := 0; i < 15; i++ {
		ch <- i
	}
	close(ch)
}

func main() {
	//var ok = true
	ch := make(chan int)

	go tel(ch)
	//for ok {
	//	if i, open := <-ch; open {
	//		fmt.Printf("ok is %t and the counter is at %d\n", ok, i)
	//	} else {
	//		break
	//	}
	//}

	// 循环打印通道数据
	for input := range ch {
		fmt.Printf("the counter is at %d\n", input)
	}
}

//the counter is at 0
//the counter is at 1
//the counter is at 2
//the counter is at 3
//the counter is at 4
//the counter is at 5
//the counter is at 6
//the counter is at 7
//the counter is at 8
//the counter is at 9
//the counter is at 10
//the counter is at 11
//the counter is at 12
//the counter is at 13
//the counter is at 14
