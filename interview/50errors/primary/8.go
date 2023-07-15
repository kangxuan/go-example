package main

import "fmt"

// var x = nil error
var x interface{} = nil // ok

func main() {
	fmt.Println(x)
}
