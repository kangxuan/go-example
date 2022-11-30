package main

import "fmt"

var barVal = map[string]int{"alpha": 34, "bravo": 56, "charlie": 23,
	"delta": 87, "echo": 56, "foxtrot": 12,
	"golf": 34, "hotel": 16, "indio": 87,
	"juliet": 65, "kili": 43, "lima": 98}

func main() {
	barValResult := make(map[int]string, len(barVal))

	fmt.Println(barVal)

	for i, v := range barVal {
		barValResult[v] = i
	}

	fmt.Println(barValResult)
}
