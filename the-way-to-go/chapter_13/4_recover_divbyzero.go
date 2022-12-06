package main

import "fmt"

// badCall1 展示被除数是0的异常
func badCall1() {
	a, b := 10, 0
	n := a / b
	fmt.Println(n)
}

func test1() {
	defer func() {
		if e := recover(); e != nil {
			fmt.Printf("Panicing %s\r\n", e)
		}
	}()
	badCall1()
	fmt.Println("After bad call")
}

func main() {
	fmt.Printf("Calling test\r\n")
	test1()
	fmt.Printf("Test completed\r\n")
}

//Calling test
//Panicing runtime error: integer divide by zero
//Test completed
