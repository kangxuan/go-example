package main

import "fmt"

func main() {
	s := []int{1, 2, 3}
	fmt.Printf("s切片的长度是：%d\n", len(s))
	fmt.Println(s)

	s = enlarge(s, 5)
	fmt.Printf("s切片的长度是：%d\n", len(s))
	fmt.Println(s)
}

func enlarge(s []int, factor int) []int {
	sl := make([]int, len(s)*factor)
	copy(sl, s)
	s = sl
	return s
}
