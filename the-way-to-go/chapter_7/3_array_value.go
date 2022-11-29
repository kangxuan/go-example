package main

import "fmt"

func main() {
	p := new([10]int)
	fmt.Println(p)

	// 数组指针
	a := [...]int{99: 1}
	var p1 = &a
	fmt.Println(p1)

	// 指针数组
	x, y := 1, 2
	a1 := [...]*int{&x, &y}
	fmt.Println(a1)

	// 因为[1]int和[2]int不是同一个类型不可用来比较
	// invalid operation: a2 == a3 (mismatched types [1]int and [2]int)
	//a2 := [1]int{}
	//a3 := [2]int{}
	//fmt.Println(a2==a3)

	a4 := [4]int{}
	a4[1] = 2
	a4[1] = 4
	fmt.Println(a4)
}
