package main

import "fmt"

func main() {
	m1 := make(map[string]string, 5)
	fmt.Println(m1)
	//fmt.Println(cap(m1)) //error
}
