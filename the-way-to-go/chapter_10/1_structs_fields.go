package main

import "fmt"

type struct1 struct {
	i1   int
	f1   float32
	str1 string
}

func main() {
	ms := new(struct1)
	ms.i1 = 10
	ms.f1 = 11.1
	ms.str1 = "Chris"

	fmt.Printf("The int is :%d\n", ms.i1)
	fmt.Printf("The float is :%f\n", ms.f1)
	fmt.Printf("The string is :%s\n", ms.str1)
	fmt.Println(ms)
}

//The int is :10
//The float is :11.100000
//The string is :Chris
//&{10 11.1 Chris}
