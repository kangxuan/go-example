package main

import "fmt"

func main() {
	s1 := "你好hello123"
	s2 := reverse([]rune(s1))
	fmt.Println(string(s2))
}

// reverse 字符串翻转
func reverse(s []rune) []rune {
	// 转换
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}
