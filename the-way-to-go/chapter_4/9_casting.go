package main

import "fmt"

func main() {
	var n int64 = 34
	var m int32
	// 将int64转换成int32，只能向精度更低的类型转换
	m = int32(n)

	fmt.Printf("32 bit int is: %d\n", m)
	fmt.Printf("64 bit int is: %d\n", n)

	a := int64(0)
	fmt.Printf("a is: %d\n", a)

	// 定义复数
	var c1 complex64 = 5 + 10i
	fmt.Println(c1)
	fmt.Printf("c1 is: %v\n", c1)
}

// 32 bit int is: 34
// 64 bit int is: 34
// a is: 0
// (5+10i)
// c1 is: (5+10i)
