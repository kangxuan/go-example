package main

import "fmt"

type Simpler interface {
	Get() int
	Set(int)
}

type Simple struct {
	i int
}

func (s Simple) Get() int {
	return s.i
}

func (s *Simple) Set(v int) {
	s.i = v
}

func main() {
	var s Simpler = &Simple{1}
	s.Set(10)
	fmt.Println("Simple i: ", s.Get())
}

//Simple i:  10
