package main

import (
	"./parse"
	"fmt"
)

func main() {
	var examples = []string{
		"1 2 3 4 5",
		"100 50 25 12.5 6.25",
		"2 + 2 = 4",
		"1st class",
		"",
	}

	for _, ex := range examples {
		fmt.Printf("Pasing %q:\n ", ex)
		nums, err := parse.Parse(ex)
		if err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Println(nums)
	}
}

//Pasing "1 2 3 4 5":
// [1 2 3 4 5]
//Pasing "100 50 25 12.5 6.25":
// pkg: pkg parse: error parsing "12.5" as int, index: 3
//Pasing "2 + 2 = 4":
// pkg: pkg parse: error parsing "+" as int, index: 1
//Pasing "1st class":
// pkg: pkg parse: error parsing "1st" as int, index: 0
//Pasing "":
// pkg: no words to parse
