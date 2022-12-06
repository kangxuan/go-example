package main

import "fmt"

func badCall() {
	panic("bad end")
}

func test() {
	// 捕获异常并恢复
	defer func() {
		if e := recover(); e != nil {
			fmt.Printf("Panicing %s\r\n", e)
		}
	}()
	badCall()
	fmt.Printf("After bad call\r\n")
}

func main() {
	fmt.Printf("Calling test\n")
	test()
	fmt.Printf("Test completed!\n")
}

//Calling test
//Panicing bad end
//Test completed!
