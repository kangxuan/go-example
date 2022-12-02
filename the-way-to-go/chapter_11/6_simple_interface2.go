package main

import "fmt"

type Simpler interface {
	Get() int
	Set(int)
}

type Simple struct {
	i int
}

func (s *Simple) Get() int {
	return s.i
}

func (s *Simple) Set(i int) {
	s.i = i
}

type RSimple struct {
	i int
	j int
}

func (p *RSimple) Get() int {
	return p.j
}

func (p *RSimple) Set(u int) {
	p.j = u
}

func fI(it Simpler) int {
	switch it.(type) {
	case *Simple:
		it.Set(5)
		return it.Get()
	case *RSimple:
		it.Set(50)
		return it.Get()
	default:
		return 99
	}
	return 0
}

func main() {
	var s Simple
	fmt.Println(fI(&s)) // &s is required because Get() is defined with a receiver type pointer
	var r RSimple
	fmt.Println(fI(&r))
}
