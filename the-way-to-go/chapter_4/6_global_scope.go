package main

var a = "A"

func main() {
	n() // A
	m() // 0
	n() // 0
}

func n() {
	print(a)
}

func m() {
	// 此处的a就是外部的a
	a = "0"
	print(a)
}
