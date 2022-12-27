package main

import (
	"fmt"
	"regexp"
)

func main() {
	a := "I am learning Go language"

	// 正则规则，2-4个字母
	re, _ := regexp.Compile("[a-z]{2,4}")

	// 查找符合正则的第一个
	one := re.Find([]byte(a))
	fmt.Println("Find:", string(one))

	// 查找符合正则的所有slice，n 小于 0 表示返回全部符合的字符串，不然就是返回指定的长度
	all := re.FindAll([]byte(a), -1)
	fmt.Println("FindAll:", all)
	for _, v := range all {
		fmt.Println(string(v))
	}

	// 查找符合条件的 index 位置, 开始位置和结束位置
	index := re.FindIndex([]byte(a))
	fmt.Println("FindIndex:", index)

	// 查找符合条件的所有的 index 位置，n 同上
	allIndex := re.FindAllIndex([]byte(a), -1)
	fmt.Println("FindAllIndex:", allIndex)

	re2, _ := regexp.Compile("am(.*)lang(.*)")
	// 查找 Submatch, 返回数组，第一个元素是匹配的全部元素，第二个元素是第一个 () 里面的，第三个是第二个 () 里面的
	subMatch := re2.FindSubmatch([]byte(a))
	for _, v := range subMatch {
		fmt.Println(string(v))
	}

	// 定义和上面的 FindIndex 一样
	subMatchIndex := re2.FindSubmatchIndex([]byte(a))
	fmt.Println(subMatchIndex)

	// FindAllSub...查找所有符合条件的子匹配
	subMatchAll := re2.FindAllSubmatch([]byte(a), -1)
	fmt.Println(subMatchAll)

	// FindAllSubmatchIndex, 查找所有字匹配的 index
	subMatchAllIndex := re2.FindAllSubmatchIndex([]byte(a), -1)
	fmt.Println(subMatchAllIndex)
}
