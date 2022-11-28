package main

import (
	"fmt"
	"strings"
)

func main() {
	var str = "a b c d "
	// Fields通过空白符分隔成slice
	var slice = strings.Fields(str)
	fmt.Printf("The Splited in slice:%v\n", slice)

	for _, val := range slice {
		fmt.Printf("%s - ", val)
	}
	fmt.Println()

	str2 := "Go1|The ABC of Go|25"
	// Split通过|分隔成slice
	slice2 := strings.Split(str2, "|")

	fmt.Printf("The Splited in slice2:%v\n", slice2)

	for _, val := range slice2 {
		fmt.Printf("%s - ", val)
	}
	fmt.Println()

	// 通过Join将slice拼接成字符串
	joinStr := strings.Join(slice2, "----")
	fmt.Printf("The Joined string is:%s\n", joinStr)
}

// The Splited in slice:[a b c d]
//a - b - c - d -
//The Splited in slice2:[Go1 The ABC of Go 25]
//Go1 - The ABC of Go - 25 -
//The Joined string is:Go1----The ABC of Go----25
