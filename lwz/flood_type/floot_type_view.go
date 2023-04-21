package main

import "fmt"

func reverseWithGenerics[T any](s []T) []T {
	l := len(s)
	r := make([]T, l)

	for i, e := range s {
		r[l-i-1] = e
	}
	return r
}

func main() {
	fmt.Println(reverseWithGenerics[int]([]int{1, 2, 3, 4}))
	fmt.Println(reverseWithGenerics[float64]([]float64{0.1, 0.2, 0.3, 0.4}))
}
