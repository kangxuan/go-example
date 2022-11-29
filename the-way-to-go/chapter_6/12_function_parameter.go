package main

func main() {
	callback(1, Add)
}

func callback(a int, f func(int, int)) {
	f(a, 2)
}

func Add(num1, num2 int) {
	println(num1 + num2)
}
