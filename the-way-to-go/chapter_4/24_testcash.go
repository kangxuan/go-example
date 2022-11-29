package main

import "fmt"

func main() {
	var p *int
	*p = 0

	fmt.Println(*p)
}
