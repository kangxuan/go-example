package main

func main() {
	res := factorial(14)
	println(res)
}

func factorial(n int) (res int) {
	if n <= 0 {
		res = 1
	} else {
		res = n * factorial(n-1)
	}

	return
}
