package main

import "fmt"

func main() {
	var num1, num2 = 1, 2
	var sum, mul, div int

	sum, mul, div = getOperationValue(num1, num2)
	fmt.Printf("%d + %d = %d\n", num1, num2, sum)
	fmt.Printf("%d * %d = %d\n", num1, num2, mul)
	fmt.Printf("%d - %d = %d\n", num1, num2, div)

	sum, mul, div = getOperationValue_2(num1, num2)
	fmt.Printf("%d + %d = %d\n", num1, num2, sum)
	fmt.Printf("%d * %d = %d\n", num1, num2, mul)
	fmt.Printf("%d - %d = %d\n", num1, num2, div)
}

// 返回值不命名
func getOperationValue(num1, num2 int) (int, int, int) {
	return num1 + num2, num1 * num2, num1 - num2
}

// 返回值命名，尽量使用命名返回值：会使代码更清晰、更简短，同时更加容易读懂。
func getOperationValue_2(num1, num2 int) (sum, mul, div int) {
	sum = num1 + num2
	mul = num1 * num2
	div = num1 - num2

	return
}
