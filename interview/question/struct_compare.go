package main

import "fmt"

type A struct {
	a int
}

type B struct {
	a int
}

func main() {
	a := A{1}
	//b := B{1}
	b := A{2}
	if a == b {
		fmt.Println("a == b")
	} else {
		fmt.Println("a != b")
	}
}
