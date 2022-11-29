package main

import "fmt"

var num = 10
var num2, num3 int

func main() {
	num2, num3 = getX2AndX3(num)
	PrintValues()
	num2, num3 = getX2AndX3_2(num)
	PrintValues()
}

func PrintValues() {
	fmt.Printf("num = %d, 2x num = %d, 3x num = %d\n", num, num2, num3)
}

func getX2AndX3(input int) (int, int) {
	num2 = 2 * input
	num3 = 3 * input

	return num2, num3
}

func getX2AndX3_2(input int) (x2, x3 int) {

	x2 = 2 * input
	x3 = 3 * input

	return

}

// num = 10, 2x num = 20, 3x num = 30
//num = 10, 2x num = 20, 3x num = 30
