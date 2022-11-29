package main

import "fmt"

func main() {
	slice1 := make([]int, 0, 20)

	caps := cap(slice1)

	for i := 0; i < caps; i++ {
		slice1 = slice1[0 : i+1]
		slice1[i] = i
		fmt.Printf("The length of slice is %d\n", len(slice1))
	}

	var num = len(slice1)
	for i := 0; i < num; i++ {
		fmt.Printf("Slice at %d is %d\n", i, slice1[i])
	}

	slice2 := slice1[1:20]
	slice2[0] = 900

	fmt.Printf("Slice1 is :%v\n", slice1)

	arr := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	a := arr[5:7]
	fmt.Printf("a : %v\n", a)

	b := a[0:5]
	fmt.Printf("b : %v\n", b)

	c := b[5:5]
	fmt.Printf("c len:%d\n", len(c))

	d := b[4:5]
	fmt.Printf("d len:%d\n", len(d))
}
