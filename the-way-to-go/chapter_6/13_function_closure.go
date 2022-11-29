package main

import "fmt"

func main() {
	f := Adder()
	fmt.Print(f(1), "-")
	fmt.Print(f(10), "-")
	fmt.Print(f(100), "-")

	var g int
	go func(i int) {
		s := 0
		for j := 0; j < i; j++ {
			s += j
		}
		g = s
		fmt.Println(g)
	}(1000)
	fmt.Println(g)
}

func Adder() func(int) int {
	var x int
	return func(delta int) int {
		x += delta
		return x
	}
}
