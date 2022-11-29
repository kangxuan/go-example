package main

import "fmt"

func main() {
	var slice1 = make([]int, 4)

	num := len(slice1)  // 长度
	caps := cap(slice1) // 容量
	fmt.Printf("slice1 len:%d, caps:%d\n", num, caps)

	// 赋值切片
	for i := 0; i < num; i++ {
		slice1[i] = i
	}

	for ix, value := range slice1 {
		fmt.Printf("Slice at %d is: %d\n", ix, value)
	}
}
