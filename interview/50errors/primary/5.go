package main

import "fmt"

func main() {
	//one := 1
	//fmt.Println(one)
	//one := 2 error：不能二次重复声明
	//fmt.Println(one)
	//var one int
	//one = 1
	//fmt.Println(one)
	one, two := work()
	fmt.Println(one)
	fmt.Println(two)
}

func work() (int, int) {
	return 1, 1
}
