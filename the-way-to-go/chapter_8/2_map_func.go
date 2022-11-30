package main

import "fmt"

func main() {
	// 证明map的值可以为任意的类型
	map1 := map[int]func() int{
		1: func() int { return 1 },
		2: func() int { return 2 },
		3: func() int { return 3 },
	}
	fmt.Println(map1)
	// map[1:0x10999e0 2:0x10999f0 3:0x1099a00]，输出了map的结构，值为函数的地址
}

//map[1:0x10999e0 2:0x10999f0 3:0x1099a00]
