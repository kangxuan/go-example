package main

import "fmt"

func main() {
	slFrom := []int{1, 2, 3}
	slTo := make([]int, 10)

	fmt.Printf("slFrom :%v, len:%d, cap:%d\n", slFrom, len(slFrom), cap(slFrom))
	fmt.Printf("slTo :%v, len:%d, cap:%d\n", slTo, len(slTo), cap(slTo))

	n := copy(slTo, slFrom)
	fmt.Printf("slTo:%v\n", slTo)
	fmt.Println(n)

	sl3 := []int{1, 2, 3}
	fmt.Printf("sl3: %v, len:%d, cap:%d\n", sl3, len(sl3), cap(sl3))
	sl3 = append(sl3, 1, 2)
	fmt.Printf("sl3: %v, len:%d, cap:%d\n", sl3, len(sl3), cap(sl3))
}
