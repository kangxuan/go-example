package main

import (
	"fmt"
	"strings"
)

func main() {
	var str = "AbcDeFG"

	fmt.Printf("The orign string is:%s\n", str)
	// 转换成小写
	fmt.Printf("The lower string is:%s\n", strings.ToLower(str))
	// 转换成大写
	fmt.Printf("The upper string is:%s\n", strings.ToUpper(str))
}

// The orign string is:AbcDeFG
//The lower string is:abcdefg
//The upper string is:ABCDEFG
