package main

import "fmt"

func main() {
	var arr1 [4]int
	f(&arr1)
}

func f(arr *[4]int) {
	fmt.Println(arr)
}
