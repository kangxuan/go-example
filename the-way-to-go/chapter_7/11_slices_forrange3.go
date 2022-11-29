package main

import "fmt"

func main() {
	items := [...]int{1, 2, 3, 4, 5}

	itemsSlice := items[:]

	for _, item := range items {
		item *= 2
	}

	for _, item := range itemsSlice {
		item *= 2
	}

	fmt.Printf("items:%v", items)

	fmt.Printf("items_slice:%v", items)
}
