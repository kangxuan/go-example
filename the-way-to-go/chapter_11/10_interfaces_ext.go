package main

import "fmt"

type Square struct {
	side float64 "边长"
}

type Triangle struct {
	base   float64 "底长"
	height float64 "高"
}

type AreaInterface interface {
	Area() float64
}

type PeriInterface interface {
	Perimeter() float64
}

func main() {
	var areaInt AreaInterface
	var periInt PeriInterface

	s := new(Square)
	s.side = 5

	t := new(Triangle)
	t.base = 3
	t.height = 4

	areaInt = s
	fmt.Println("The square has area: ", areaInt.Area())

	periInt = s
	fmt.Printf("The square has perimeter: %f\n", periInt.Perimeter())

	areaInt = t
	fmt.Println("The Triangle has peri: ", areaInt.Area())
}

func (s *Square) Area() float64 {
	return s.side * s.side
}

func (s *Square) Perimeter() float64 {
	return 4 * s.side
}

func (t *Triangle) Area() float64 {
	return t.base * t.height / 2
}

//The square has area:  25
//The square has perimeter: 20.000000
//The Triangle has peri:  6
