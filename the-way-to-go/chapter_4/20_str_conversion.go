package main

import (
	"fmt"
	"strconv"
)

func main() {
	var orignStr = "666"
	var an int
	var newS string

	fmt.Printf("The size of ints is:%d\n", strconv.IntSize)

	// 将字符串转成整型
	an, _ = strconv.Atoi(orignStr)
	fmt.Printf("The integer is:%d\n", an)

	an += 5

	// 将整型转换成字符串
	newS = strconv.Itoa(an)
	fmt.Printf("The new string is:%s\n", newS)
}

// The size of ints is:64
//The integer is:666
//The new string is:671
