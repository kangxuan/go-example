package main

import (
	"fmt"
	"strconv"
)

type TwoInts struct {
	a int
	b int
}

func (tn *TwoInts) String() string {
	return "(" + strconv.Itoa(tn.a) + "/" + strconv.Itoa(tn.b) + ")"
}

func main() {
	t1 := new(TwoInts)
	t1.a = 3
	t1.b = 4

	fmt.Printf("two1 is: %v\n", t1)
	fmt.Println("two1 is:", t1)
	fmt.Printf("two1 is: %T\n", t1)
	fmt.Printf("two1 is: %#v\n", t1)
}

//two1 is: (3/4)
//two1 is: (3/4)
//two1 is: *main.TwoInts
//two1 is: &main.TwoInts{a:3, b:4}
