package main

import "fmt"

func main() {
	var first = 10
	var crond int

	fmt.Printf("first : %d, crond : %d\n", first, crond)

	if first <= 0 {
		fmt.Printf("first <= 0\n")
	} else if first > 0 && first <= 5 {
		fmt.Printf("fist > 0 && first <= 5\n")
	} else {
		fmt.Printf("first > 5\n")
	}

	if crond := 10; crond > 10 {
		fmt.Printf("crond > 10\n")
	} else {
		fmt.Printf("crond <= 10\n")
	}

	fmt.Printf("crond : %d", crond)
}
