package main

import (
	"fmt"
	"strings"
)

func main() {
	var str = "Hi, I'm Marc, Hi."

	fmt.Printf("The position of \"Marc\" is: ")
	// Index() 返回字符串 str 在字符串 s 中的索引（str 的第一个字符的索引），-1 表示字符串 s 不包含字符串 str
	fmt.Printf("%d\n", strings.Index(str, "Marc"))

	fmt.Printf("The position of the first instance of \"Hi\" is: ")
	fmt.Printf("%d\n", strings.Index(str, "Hi"))
	fmt.Printf("The position of the last instance of \"Hi\" is: ")
	fmt.Printf("%d\n", strings.LastIndex(str, "Hi"))

	fmt.Printf("The position of \"Burger\" is: ")
	fmt.Printf("%d\n", strings.Index(str, "Burger"))

	var str1 = "你好，中国"
	fmt.Printf("The position of \"中国\": ")
	fmt.Printf("%d\n", strings.IndexRune(str1, rune('中')))
}

// The position of "Marc" is: 8
//The position of the first instance of "Hi" is: 0
//The position of the last instance of "Hi" is: 14
//The position of "Burger" is: -1
//The position of "中国": 9
