package main

import "fmt"

type Shaper interface {
	Area() float32
}

type Square struct {
	side float32
}

func (sq *Square) Area() float32 {
	return sq.side * sq.side
}

type Rectangle struct {
	lenght, width float32
}

func (r Rectangle) Area() float32 {
	return r.lenght * r.width
}

func main() {
	r := Rectangle{5, 3}
	q := &Square{5}

	shapers := []Shaper{r, q}
	fmt.Println("Looping through shapes for area ...")
	for n := range shapers {
		fmt.Println("Shaper details: ", shapers[n])
		fmt.Println("Area of this shaper is: ", shapers[n].Area())
	}
}

//Looping through shapes for area ...
//Shaper details:  {5 3}
//Area of this shaper is:  15
//Shaper details:  &{5}
//Area of this shaper is:  25
