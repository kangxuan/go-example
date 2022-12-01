package main

import "fmt"

type twoInt struct {
	a, b int
}

func (t *twoInt) addThem() int {
	t.a++
	return t.a + t.b
}

func (t *twoInt) addToParam(param int) int {
	return t.a + t.b + param
}

func main() {
	t := &twoInt{1, 2}
	s1 := t.addThem()
	fmt.Println(s1)
	s2 := t.addToParam(1)
	fmt.Println(s2)
}

//4
//5
