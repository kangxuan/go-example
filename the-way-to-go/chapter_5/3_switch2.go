package main

import (
	"fmt"
)

func main() {
	var num2 = 10

	switch {
	case num2 == 10:
		fmt.Println("中了")
		fallthrough
	default:
		fmt.Println("差点意思")
	}
}
