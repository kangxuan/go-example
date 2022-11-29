package main

import (
	"fmt"
	"strings"
)

func main() {

	var str = " fasdf  "
	fmt.Printf("The origin string is:%s\n", str)
	fmt.Printf("The trimSpace string is:%s\n", strings.TrimSpace(str))
	fmt.Printf("The trim string is:%s\n", strings.Trim(str, " "))
	fmt.Printf("The trimLeft string is:%s\n", strings.TrimLeft(str, " "))
	fmt.Printf("The trimRight string is:%s\n", strings.TrimRight(str, " "))
}
