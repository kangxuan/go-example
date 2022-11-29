package main

func main() {
	n := 0
	// 引用
	reply := &n

	Miltiply(10, 5, reply)

	print(*reply)
}

func Miltiply(a, b int, reply *int) {
	*reply = a * b
}
