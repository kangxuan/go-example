package main

import "fmt"

func main() {
	f()
}

func f() {
	for i := 0; i < 4; i++ {
		g := func(i int) {
			fmt.Printf("%d ", i)
		}
		g(i)
		fmt.Printf(" - g is of type %T and has value %v\n", g, g)
	}
}

//0  - g is of type func(int) and has value 0x109af60
//1  - g is of type func(int) and has value 0x109af60
//2  - g is of type func(int) and has value 0x109af60
//3  - g is of type func(int) and has value 0x109af60
