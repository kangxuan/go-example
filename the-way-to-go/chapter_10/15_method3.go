package main

import "fmt"
import "math"

type Point struct {
	x, y float64
}

func (p *Point) abs() float64 {
	return math.Sqrt(p.x*p.x + p.y*p.y)
}

type NamedPoint struct {
	Point
	name string
}

// 可以理解为对Point的abs方法的重载
func (np *NamedPoint) abs() float64 {
	return np.Point.abs() * 100
}

func main() {
	n := &NamedPoint{Point{3, 4}, "kangxuan"}
	// 出现重载方法则需要具体调用方法
	fmt.Println(n.Point.abs())
	fmt.Println(n.abs())
}

//5
//500
