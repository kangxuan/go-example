package main

import "fmt"

func main() {
	s := []int{0, 1, 2, 3, 4, 5, 6}

	s = Filter(s, even)

	fmt.Println(s)
}

// 通过传函数返回符合条件的新切片
func Filter(s []int, fn func(int) bool) []int {
	var sl []int
	for _, value := range s {
		if fn(value) {
			sl = append(sl, value)
		}
	}
	return sl
}

// 偶数返回真
func even(n int) bool {
	if n%2 == 0 {
		return true
	}
	return false
}
