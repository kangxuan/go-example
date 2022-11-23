package main

func main() {
	a, b := 15, 0
	// 0不能做除数
	c := a / b
	print(c)
}

// panic: runtime error: integer divide by zero
