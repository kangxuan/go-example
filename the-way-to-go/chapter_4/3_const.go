package main

import "fmt"

// 显式类型定义
const T string = "1234"

// 隐式类型定义，会自行推断类型
const Pi = 3.1415

// 多个常量同时定义
const Monday, Tuesday, Wednesday, Thursday, Friday, Saturday = 1, 2, 3, 4, 5, 6

// 常量组定义嗯
const (
	B float64 = 1 << (iota * 10) // iota可以被用作枚举值，第一个为0，依次递增，在每遇到一个新的常量块或单个常量声明时， iota 都会重置为 0
	KB
	MB
)

// 因为在编译期间自定义函数均属于未知，因此无法用于常量的赋值，会编译报错
//const I = getNum()

func main() {
	fmt.Println(T)
	fmt.Println(Pi)
	fmt.Println("Monday:", Monday, "Tuesday:", Tuesday, "Wednesday:", Wednesday, "Thursday:", Thursday, "Friday:", Friday, "Saturday:", Saturday)
	fmt.Println("B:", B, "KB:", KB, "MB:", MB)
	//fmt.Println(I)
}

//func getNum() int {
//	return 1
//}

// 1234
// 3.1415
// Monday: 1 Tuesday: 2 Wednesday: 3 Thursday: 4 Friday: 5 Saturday: 6
// B: 1 KB: 1024 MB: 1.048576e+06
