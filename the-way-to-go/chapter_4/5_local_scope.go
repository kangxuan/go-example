package main

var a = "G"

func main() {
	n() // G
	m() // 0
	n() // G
}

// 作用域
func n() {
	print(a)
}

func m() {
	// 在m中的a和外部的a不是同一个数据
	a := "0"
	print(a)
}
