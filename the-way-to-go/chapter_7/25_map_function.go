package main

import "fmt"

func main() {
	list := []int{1, 2, 3, 4}
	// 定义个函数
	mf := func(i int) int {
		return i * 10
	}

	fmt.Println(list)
	result := mapFunc(mf, list)
	fmt.Println(result)
}

func mapFunc(mf func(int) int, list []int) []int {
	result := make([]int, len(list))

	for i, value := range list {
		result[i] = mf(value)
	}

	return result
}
