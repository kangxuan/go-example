package main

var a string

func main() {
	a = "G"
	print(a) // G
	f1()
}

func f1() {
	a := "0"
	print(a) // 0
	f2()
}

func f2() {
	print(a) // G
}
