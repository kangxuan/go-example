package main

import "fmt"

func main() {
	sl := []int{1, 2, 5, 2, 3, 5}

	s := bubbleSort(sl)
	fmt.Println(s)
}

// 冒泡排序
func bubbleSort(sl []int) []int {
	sLen := len(sl)
	var tmp int
	for i := 0; i < sLen; i++ {
		for j := 0; j < sLen; j++ {
			if sl[i] > sl[j] {
				tmp = sl[i]
				sl[i] = sl[j]
				sl[j] = tmp
			}
		}
	}
	return sl
}
