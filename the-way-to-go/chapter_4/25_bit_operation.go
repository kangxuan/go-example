package main

import "fmt"

func main() {
	var x uint8 = 15
	var y uint8 = 4

	fmt.Printf("%08b\n", x&^y)

	fmt.Printf("1 & 1: %d\n", 1&1)
	fmt.Printf("1 | 1: %d\n", 1|1)
	fmt.Printf("1 ^ 1: %d\n", 1^1)
	fmt.Printf("1 &^ 1: %d\n", 1&^1)

	//var z float32 = 3
	//fmt.Printf("%d\n", x&z)
	//报错：invalid operation: x & z (mismatched types uint8 and float32)
}

// 00001011
//1 & 1: 1
//1 | 1: 1
//1 ^ 1: 0
//1 &^ 1: 0
