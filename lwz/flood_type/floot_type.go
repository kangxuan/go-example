package main

import "fmt"

func reverse(s []int) []int {
	l := len(s)
	r := make([]int, l)

	for i, e := range s {
		r[l-i-1] = e
	}
	return r
}

func reverseFloat64Slice(s []float64) []float64 {
	l := len(s)
	r := make([]float64, l)

	for i, e := range s {
		r[l-i-1] = e
	}
	return r
}

func main() {
	fmt.Println(reverse([]int{1, 2, 3, 4}))
	fmt.Println(reverseFloat64Slice([]float64{0.1, 0, 2, 0.3, 0.4}))
}
