package main

import (
	"fmt"
)

func main() {
	var num1 = 10
	switch num1 {
	case 1:
		fmt.Println("is 1")
	case 2:
		fmt.Println("is 2")
	case 10:
		fmt.Println("is 10")
	default:
		fmt.Println("other")
	}
}
