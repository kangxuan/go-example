package main

import (
	"fmt"
	"time"
)

const LIMIT = 25

// map
var fibs [LIMIT + 1]uint64

func main() {
	start := time.Now()
	var res uint64
	for i := 1; i <= LIMIT; i++ {
		res = fibonacci(i)
		println(res)
	}

	end := time.Now()
	delta := end.Sub(start)
	fmt.Printf("执行时间：%s\n", delta)

}

func fibonacci(i int) (res uint64) {
	if fibs[i] != 0 {
		res = fibs[i]
		return
	}
	if i <= 1 {
		res = 1
	} else {
		res = fibonacci(i-1) + fibonacci(i-2)
	}

	fibs[i] = res

	return
}
