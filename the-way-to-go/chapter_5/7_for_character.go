package main

import "fmt"

func main() {
	for i := 0; i < 25; i++ {
		for j := 0; j < i; j++ {
			print("G")
		}
		fmt.Println()
	}

	str := "G"
	for i := 0; i < 25; i++ {
		fmt.Println(str)
		str += "G"
	}
}
