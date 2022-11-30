package main

import "fmt"

func main() {
	// 创建一个切片，类型为map，长度为5，cap为5
	items := make([]map[int]int, 5)
	// 循环储值map
	for i := range items {
		items[i] = make(map[int]int, 1)
		items[i][1] = 2
	}

	fmt.Printf("The items is:%v\n", items)

	// 创建一个map切片，实际上是无法修改items2
	items2 := make([]map[int]int, 5)
	for _, v := range items2 {
		v = make(map[int]int, 1)
		v[1] = 3
	}
	fmt.Printf("The items is:%v\n", items2)
}

//The items is:[map[1:2] map[1:2] map[1:2] map[1:2] map[1:2]]
//The items is:[map[] map[] map[] map[] map[]]
