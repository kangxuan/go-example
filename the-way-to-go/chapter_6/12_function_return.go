package main

import "fmt"

func main() {
	f := Add()
	fmt.Printf("f(2): %d\n", f(2))

	f2 := Adder(2)
	fmt.Printf("f2(3): %d\n", f2(3))
}

// 将匿名函数作为返回值返回
func Add() func(b int) int {
	return func(b int) int {
		return b + 2
	}
}

func Adder(a int) func(b int) int {
	return func(b int) int {
		return a + b
	}
}
