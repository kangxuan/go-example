package main

import "fmt"

func main() {
	a, b, c := 1, 2, 3
	fmt.Printf("乘法：%d * %d * %d = %d\n", a, b, c, MultiPly3Nums(a, b, c))
}

func MultiPly3Nums(a, b, c int) int {
	return a * b * c
}
