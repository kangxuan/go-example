package main

import "fmt"

// 定义一个矩形结构体
type Rectangle struct {
	length, width float64
}

// 定义求面积的方法
func (r *Rectangle) Area() float64 {
	a := r.length * r.width
	return a
}

// 定义求周长的方法
func (r *Rectangle) Perimeter() float64 {
	p := 2 * (r.length * r.width)
	return p
}

func main() {
	p1 := &Rectangle{12, 32}
	fmt.Printf("p1.length : %f, width : %f, area: %f, perimeter: %f\n", p1.length, p1.width, p1.Area(), p1.Perimeter())
}

//p1.length : 12.000000, width : 32.000000, area: 384.000000, perimeter: 768.000000
