package main

import "fmt"

func main() {
	var slice1 = make([]int, 10)
	for i := 0; i < len(slice1); i++ {
		slice1[i] = i * 10
	}

	for i := 0; i < len(slice1); i++ {
		fmt.Printf("Slice1 at %d is %d\n", i, slice1[i])
	}

	fmt.Printf("\nThe length of slice1 is %d\n", len(slice1))
	fmt.Printf("The capacity of slice1 is %d\n", cap(slice1))

	slice2 := make([]byte, 5)
	fmt.Printf("The length of slice2 is %d\n", len(slice2))
	fmt.Printf("The cap of slice2 is %d\n", cap(slice2))
	slice2 = slice2[2:4]
	fmt.Printf("The length of slice[2:4] is %d\n", len(slice2))
	fmt.Printf("The cap of slice[2:4] is %d\n", cap(slice2))

	s1 := []byte{'p', 'o', 'e', 'm'}
	s2 := s1[2:]
	fmt.Printf("s1:%s，s2:%s\n", s1, s2)

	s2[1] = 't'
	fmt.Printf("s1:%s，s2:%s\n", s1, s2)

	var arr1 = [3]int{1, 2, 3}
	s3 := arr1[1:2]
	fmt.Printf("s3:%v\n", s3)
	s3[0] = 11
	fmt.Printf("arr1:%v\n", arr1)
}
