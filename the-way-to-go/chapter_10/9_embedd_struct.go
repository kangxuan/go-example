package main

import "fmt"

type A struct {
	ax, ay int
}

type B struct {
	A
	bx, by float32
}

type C struct {
	A
	ax, ay int
}

func main() {
	b := B{A{1, 2}, 3.0, 5.}
	fmt.Println(b)

	c := C{A{1, 2}, 3, 4}
	fmt.Println(c)

	c1 := &C{}
	c1.ax = 4
	c1.ay = 5
	c1.A.ax = 1
	c1.A.ay = 2
	fmt.Println(c1)
}
