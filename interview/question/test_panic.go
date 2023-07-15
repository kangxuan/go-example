package main

func main() {
	//var x interface{}
	//var y interface{} = []int{3, 5}
	//_ = x == x
	//_ = x == y
	//_ = y == y

	x := make([]int, 2, 10)
	_ = x[6:10]
	//_ = x[6:]
	_ = x[2:]
}
