package main

import "fmt"

func main() {
	str := "aabbb"
	pos := len(str) / 2
	newStr := str[pos:] + str[:pos]
	fmt.Println(newStr)
}
