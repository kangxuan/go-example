package main

import "fmt"

func main() {
	printInts()
	fmt.Println()
	printInts(2, 3)
	fmt.Println()
	printInts([]int{1, 2, 3, 4, 5}...)
}

// 变长参数，如果变长参数不是同一个类型有两种解决方案，1.采用结构体，2.空接口
func printInts(list ...int) {
	for _, v := range list {
		fmt.Println(v)
	}
}
