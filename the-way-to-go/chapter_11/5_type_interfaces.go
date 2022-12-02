package main

import (
	"fmt"
	"math"
)

type Shaper interface {
	Area() float32
}

type Square struct {
	side float32
}

type Circle struct {
	radius float32
}

func (s Square) Area() float32 {
	return s.side * s.side
}

func (c Circle) Area() float32 {
	return c.radius * c.radius * math.Pi
}

func main() {
	var areaIntf Shaper

	sq1 := new(Square)
	sq1.side = 10

	areaIntf = *sq1

	if t, ok := areaIntf.(Square); ok {
		fmt.Printf("The type of areaIntf is: %T\n", t)
	}

	ci1 := Circle{10}
	areaIntf = ci1

	if u, ok := areaIntf.(Circle); ok {
		fmt.Printf("The type of areaIntf is: %T\n", u)
	} else {
		fmt.Println("areaIntf does not contain a variable of type Circle")
	}

	if t, ok := areaIntf.(Square); ok {
		fmt.Printf("The type of areaIntf is: %T\n", t)
	}
}

//The type of areaIntf is: main.Square
//The type of areaIntf is: main.Circle
