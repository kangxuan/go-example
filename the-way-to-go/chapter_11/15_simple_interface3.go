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

func gI(any interface{}) int {
	if v, ok := any.(Simpler); ok {
		return v.Get()
	}
	return 0
}

func main() {
	var s Simple
	fmt.Println(fI(&s)) // &s is required because Get() is defined with a receiver type pointer
	var r RSimple
	fmt.Println(fI(&r))

	var s1 = Simple{6}
	fmt.Println(gI(&s1)) // &s is required because Get() is defined with a receiver type pointer
	var r1 = RSimple{60, 60}
	fmt.Println(gI(&r1))
}

//5
//50
//6
//60
