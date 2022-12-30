package main

import (
	"fmt"
	"strconv"
)

// 展示strconv相关操作
func main() {
	// Append 系列函数将整数等转换为字符串后，添加到现有的字节数组中。
	str := make([]byte, 0, 100)
	str = strconv.AppendInt(str, 4567, 10)
	str = strconv.AppendBool(str, false)
	str = strconv.AppendQuote(str, "abcdefg")
	str = strconv.AppendQuoteRune(str, '单')
	fmt.Println(string(str))

	// Format 系列函数把其他类型的转换为字符串
	a := strconv.FormatBool(false)
	b := strconv.FormatInt(4567, 10)
	c := strconv.FormatFloat(123.23, 'g', 12, 64)
	d := strconv.FormatUint(12345, 10)
	e := strconv.Itoa(1023)
	fmt.Println(a, b, c, d, e)

	// Parse 系列函数把字符串转换为其他类型
	f, err := strconv.ParseBool("false")
	checkError(err)
	g, err := strconv.ParseFloat("123.23", 64)
	checkError(err)
	h, err := strconv.ParseInt("1234", 10, 64)
	checkError(err)
	i, err := strconv.ParseUint("123456", 10, 64)
	checkError(err)
	j, err := strconv.Atoi("1024")
	fmt.Println(f, g, h, i, j)
}

func checkError(e error) {
	if e != nil {
		fmt.Println(e)
	}
}
