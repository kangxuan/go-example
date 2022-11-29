package main

import "fmt"

func main() {
	A()
	B()
	C()
}

func A() {
	fmt.Println("Func A")
}

func B() {
	// defer 通过recover对panic进行捕获
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Recover in B")
		}
	}()
	panic("panic B")
}

func C() {
	fmt.Println("Func C")
}

// Func A
//Recover in B
//Func C
