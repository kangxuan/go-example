package main

import "fmt"

func main() {
	b()
}

func trace(s string) string {
	fmt.Println("Entering:", s)
	return s
}

func untrace(s string) string {
	fmt.Println("Leaving:", s)
	return s
}

func a() {
	defer untrace(trace("a"))
	fmt.Println("in a")
}

func b() {
	defer untrace(trace("b"))
	fmt.Println("in b")
	a()
}
