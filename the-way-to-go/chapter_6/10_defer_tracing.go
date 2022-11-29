package main

import "fmt"

func main() {
	b()
}

func trace(s string) {
	fmt.Println("Entering:", s)
}

func untrace(s string) {
	fmt.Println("Leaving:", s)
}

func a() {
	trace("a")
	defer untrace("a")
	fmt.Println("in a")
}

func b() {
	trace("b")
	defer untrace("b")
	fmt.Println("in b")
	a()
}

//Entering: b
//in b
//Entering: a
//in a
//Leaving: a
//Leaving: b
