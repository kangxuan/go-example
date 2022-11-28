package main

import (
	"fmt"
	"strings"
)

func main() {
	var str = "KKKKKKKK is XXXXK"
	var str1 = "GGGGGGG"

	// Count() 用于计算字符串 str 在字符串 s 中出现的非重叠次数
	fmt.Printf("str \":%s\" K's count is:%d\n", str, strings.Count(str, "K"))
	fmt.Printf("str1 \":%s\" GG's count is:%d\n", str1, strings.Count(str1, "GG"))
}

// str ":KKKKKKKK is XXXXK" K's count is:9
//str1 ":GGGGGGG" GG's count is:3
