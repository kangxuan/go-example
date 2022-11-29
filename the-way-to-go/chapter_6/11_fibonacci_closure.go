package main

import (
	"fmt"
	"time"
)

func main() {
	f := fib()

	// 统计函数执行时间
	start := time.Now()
	for i := 0; i <= 9; i++ {
		println(i+2, f())
	}
	end := time.Now()
	delta := end.Sub(start)
	fmt.Printf("longCalculation took this amount of time: %s\n", delta)
}

func fib() func() int {
	a, b := 1, 1
	return func() int {
		a, b = b, a+b
		return b
	}
}
