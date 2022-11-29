package main

import "fmt"

func main() {
	var n = 3
	b := fibonacciFunc(n)
	fmt.Printf("%v", b)
}

func fibonacciFunc(n int) (res []int) {
	if n <= 1 {
		res[n] = 1
	} else {
		res[n] = fibonacciFunc(n - 1)[n-1] + fibonacciFunc(n - 2)[n-2]
	}

	return
}
