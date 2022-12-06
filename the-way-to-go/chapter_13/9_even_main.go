package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 5; i++ {
		fmt.Printf("Is the integer %d even? %v\n", i, Even(i))
	}
}

// Even 判断偶数
func Even(i int) bool {
	return i%2 == 0
}

//Is the integer 0 even? true
//Is the integer 1 even? false
//Is the integer 2 even? true
//Is the integer 3 even? false
//Is the integer 4 even? true
