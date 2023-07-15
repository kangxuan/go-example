package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	x := []string{"a", "b", "c"}
	for v := range x {
		fmt.Println(v)
	}
	s4 := make([]int, 0)
	fmt.Printf("s4 pointer:%+v, \n", *(*reflect.SliceHeader)(unsafe.Pointer(&s4)))
}
