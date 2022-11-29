package main

import "fmt"

func main() {
	pos := 10
	res, pos := fibonacci(pos)
	fmt.Printf("pos:%d, val:%d\n", pos, res)
}

func fibonacci(n int) (res, pos int) {
	if n <= 1 {
		res = 1
	} else {
		res1, _ := fibonacci(n - 1)
		res2, _ := fibonacci(n - 2)
		res = res1 + res2
	}

	pos = n

	return
}
