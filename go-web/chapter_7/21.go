package main

import (
	"fmt"
	"strings"
)

// 展示字符串操作,strings包
func main() {
	// 字符串 s 中是否包含 substr，返回 bool 值
	fmt.Println(strings.Contains("seafood", "foo"))
	fmt.Println(strings.Contains("seafood", "bar"))
	fmt.Println(strings.Contains("seafood", ""))
	fmt.Println(strings.Contains("", ""))

	// 字符串链接，把 slice a 通过 sep 链接起来
	s := []string{"foo", "bar", "baz"}
	fmt.Println(strings.Join(s, ", "))

	// 字符串分隔
	s1 := "foo,bar,baz"
	fmt.Println(strings.Split(s1, ","))

	// 在字符串 s 中查找 sep 所在的位置，返回位置值，找不到返回 -1
	fmt.Println(strings.Index("chicken", "ken"))
	fmt.Println(strings.Index("chicken", "dm"))

	// 重复 s 字符串 count 次，最后返回重复的字符串
	fmt.Println("ba" + strings.Repeat("na", 2))

	// 在 s 字符串中，把 old 字符串替换为 new 字符串，n 表示替换的次数，小于 0 表示全部替换
	fmt.Println(strings.Replace("oink oink oink", "k", "ky", 2))
	fmt.Println(strings.Replace("oink oink oink", "oink", "moo", -1))

	// 在 s 字符串的头部和尾部去除 cutset 指定的字符串
	fmt.Println(strings.Trim(" !!!Achtung !!! ", "! "))

	// 去除 s 字符串的空格符，并且按照空格分割返回 slice
	fmt.Printf("Fields are : %q", strings.Fields(" foo bar baz    "))
}

//true
//false
//true
//true
//foo, bar, baz
//[foo bar baz]
//4
//-1
//banana
//oinky oinky oink
//moo moo moo
//Achtung
//Fields are : ["foo" "bar" "baz"]
