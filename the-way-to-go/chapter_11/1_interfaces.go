package main

import "fmt"

// 定义一个Shaper接口
type Shaper interface {
	Area() float32
}

// 定义一个正方形结构
type Square struct {
	side float32
}

// 实现Area()方法
func (sq *Square) Area() float32 {
	return sq.side * sq.side
}

func main() {
	sq1 := new(Square)
	sq1.side = 5

	var areaIntf Shaper
	// 将sq1赋值给areaIntf
	areaIntf = sq1

	fmt.Printf("The square has area: %f\n", areaIntf.Area())
}

//The square has area: 25.000000
