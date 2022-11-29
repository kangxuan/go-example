package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	var res int
	for i := 0; i <= 45; i++ {
		res = fibonacci(i)
		fmt.Printf("fibonacci(%d) is: %d\n", i, res)
	}

	end := time.Now()
	delta := end.Sub(start)
	fmt.Printf("执行时间：%s\n", delta)

}

func fibonacci(i int) (res int) {
	if i <= 1 {
		res = 1
	} else {
		res = fibonacci(i-1) + fibonacci(i-2)
	}

	return
}
