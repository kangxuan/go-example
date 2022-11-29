package main

import "fmt"

func main() {
	str := "GO语言编程"

	fmt.Println(len(str))
	fmt.Println(len([]rune(str)))

	fmt.Println(str[0:8])
	fmt.Println(str[0:7])
	fmt.Println(string([]rune(str)[:4]))
}
