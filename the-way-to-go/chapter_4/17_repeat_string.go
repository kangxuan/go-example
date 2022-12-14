package main

import (
	"fmt"
	"strings"
)

func main() {
	var origS = "Hi there!\n"
	var newS string

	// Repeat() 用于重复 count 次字符串 s 并返回一个新的字符串：
	newS = strings.Repeat(origS, 3)
	fmt.Printf("The new repeated string is:%s\n", newS)
}

// The new repeated string is:Hi there!
//Hi there!
//Hi there!
