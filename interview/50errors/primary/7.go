package main

import "fmt"

func main() {
	x := 1
	fmt.Println(x)
	{
		fmt.Println(x)
		x = 3
		fmt.Println(x)
	}
	fmt.Println(x)
}
