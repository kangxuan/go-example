package main

import "fmt"

func main() {
	function1()
	a()
	f()
}

func function1() {
	fmt.Println("In function1 at the top")
	defer function2()
	function2()
	fmt.Printf("In function1 at the bottom!\n")
}

func function2() {
	fmt.Printf("Function2: Deferred until the end of the calling function!\n")
}

func a() {
	i := 0
	i++
	defer fmt.Println(i)
	i = 4000
	return
}

func f() {
	// 当有多个 defer 行为被注册时，它们会以逆序执行（类似栈，即后进先出）
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}

// In function1 at the top
//Function2: Deferred until the end of the calling function!
//In function1 at the bottom!
//Function2: Deferred until the end of the calling function!
//1
//4
//3
//2
//1
//0
