package main

import "fmt"

func main() {
	// 声明一个数组
	var arr [5]int

	// 给数字赋值
	for i := 0; i < 5; i++ {
		arr[i] = i * 10
	}

	// 输出打印
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}

	var arr1 = [...]int{1, 2, 3, 4, 5}

	for i := range arr1 {
		fmt.Println(arr1[i])
	}
}

//0
//10
//20
//30
//40
//1
//2
//3
//4
//5
