package main

import "fmt"

// type一个int的切片
type IntVector []int

func (iv IntVector) Sum() (s int) {
	for _, v := range iv {
		s += v
	}
	return
}

func main() {
	v := IntVector{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(v.Sum())
}

//55
