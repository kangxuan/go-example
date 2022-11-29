package main

import "fmt"

func main() {
	// 定义一个数组
	var arr1 [6]int
	// 定义一个切片
	var slice1 = arr1[2:5]

	fmt.Printf("slice1的长度：%d, 容量：%d，%v\n", len(slice1), cap(slice1), slice1)

	// 给arr1数组赋值
	for i := 0; i < len(arr1); i++ {
		arr1[i] = i
	}

	for i := 0; i < len(slice1); i++ {
		fmt.Printf("Slice at %d is %d\n", i, slice1[i])
	}

	fmt.Printf("The length of arr1 is %d\n", len(arr1))
	fmt.Printf("The length of slice1 is %d\n", len(slice1))
	fmt.Printf("The capacity of slice1 is %d\n", cap(slice1))

	slice1 = slice1[0:4]
	for i := 0; i < len(slice1); i++ {
		fmt.Printf("Slice at %d is %d\n", i, slice1[i])
	}

	fmt.Printf("The length of slice1 is %d\n", len(slice1))
	fmt.Printf("The capacity of slice1 is %d\n", cap(slice1))

	b := []string{"g", "o", "l", "a", "n", "g"}
	fmt.Printf("b[1:4]：%v\n", b[1:4])
	fmt.Printf("b[:2]：%v\n", b[:2])
	fmt.Printf("b[2:]：%v\n", b[2:])
	fmt.Printf("b[:]：%v\n", b[:])
}

//slice1的长度：3, 容量：4，[0 0 0]
//Slice at 0 is 2
//Slice at 1 is 3
//Slice at 2 is 4
//The length of arr1 is 6
//The length of slice1 is 3
//The capacity of slice1 is 4
//Slice at 0 is 2
//Slice at 1 is 3
//Slice at 2 is 4
//Slice at 3 is 5
//The length of slice1 is 4
//The capacity of slice1 is 4
//b[1:4]：[o l a]
//b[:2]：[g o]
//b[2:]：[l a n g]
//b[:]：[g o l a n g]
