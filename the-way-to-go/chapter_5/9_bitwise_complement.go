package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Printf("i的二进制数是: %b，按位补齐的结果是：%b\n", i, ^i)
	}
}
