package main

import "fmt"

type number struct {
	f float32
}

type nr = number

func main() {
	var n1 = number{1.0}
	var n2 = nr{2.0}
	n1 = n2
	fmt.Println(n1, n2)
}
