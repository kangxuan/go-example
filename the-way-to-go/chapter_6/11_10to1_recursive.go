package main

import "fmt"

func main() {
	printNum(10)
}

func printNum(n int) {
	if n <= 0 {
		fmt.Println("打印结束")
	} else {
		fmt.Println(n)
		printNum(n - 1)
	}
}
