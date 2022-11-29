package main

import "fmt"

func main() {
	x := min(1, 1, 2, 0, -1)
	fmt.Println(x)

	slice := []int{1, 1, 2, -2, -1}
	x = min(slice...)
	fmt.Println(x)
}

// 变长参数
func min(s ...int) int {
	if len(s) <= 0 {
		return 0
	}

	min := s[0]
	for _, v := range s {
		if v < min {
			min = v
		}
	}

	return min
}
