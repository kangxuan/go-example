package main

import "fmt"

type employee struct {
	name   string
	gender int
	salary float32
}

func (e *employee) AddSalary(ratio float32) {
	e.salary += e.salary * (1 + ratio)
}

func main() {
	emp1 := &employee{"康宣", 1, 90000}
	emp1.AddSalary(0.5)
	fmt.Println(*emp1)
}
