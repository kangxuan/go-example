package main

import "fmt"

func main() {
	seasons := []string{"Spring", "Summer", "Autumn", "Winter"}

	season1 := seasons[1:3]

	fmt.Printf("seasons len:%d, cap:%d\n", len(seasons), cap(seasons))

	fmt.Printf("season1 : %v\n", season1)

	fmt.Printf("seasons type:%T, season1 type:%T\n", seasons, season1)

	for ix, season := range seasons {
		fmt.Printf("Season %d is: %s\n", ix, season)
	}

	var season string
	for _, season = range seasons {
		fmt.Printf("%s\n", season)
	}
}
