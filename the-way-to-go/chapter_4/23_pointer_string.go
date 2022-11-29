package main

import "fmt"

//const i = 5

func main() {
	s := "good bye"
	var p = &s
	*p = "ciao"

	fmt.Printf("Here is the pointer p: %p\n", p)
	fmt.Printf("Here is the string *p:%s\n", *p)
	fmt.Printf("Here is the string s:%s\n", s)
}

// Here is the pointer p: 0xc0000721e0
//Here is the string *p:ciao
//Here is the string s:ciao
