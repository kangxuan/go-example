package main

import (
	"fmt"
	"sort"
)

func main() {
	drinks := map[string]string{"beer": "啤酒",
		"wine":   "葡萄酒",
		"water":  "白开水",
		"coffee": "咖啡",
		"tea":    "茶"}

	sDrinks := make([]string, len(drinks))

	ix := 0

	fmt.Printf("The following drinks are available:\n")
	for eng := range drinks {
		sDrinks[ix] = eng
		ix++
		fmt.Println(eng)
	}

	fmt.Println("")
	for eng, fr := range drinks {
		fmt.Printf("The french for %s is %s\n", eng, fr)
	}

	fmt.Println("")
	fmt.Println("Now the sorted output:")
	sort.Strings(sDrinks)

	fmt.Printf("The following sorted drinks are available:\n")
	for _, eng := range sDrinks {
		fmt.Println(eng)
	}

	fmt.Println("")
	for _, eng := range sDrinks {
		fmt.Printf("The french for %s is %s\n", eng, drinks[eng])
	}
}

//The following drinks are available:
//tea
//beer
//wine
//water
//coffee
//
//The french for beer is 啤酒
//The french for wine is 葡萄酒
//The french for water is 白开水
//The french for coffee is 咖啡
//The french for tea is 茶
//
//Now the sorted output:
//The following sorted drinks are available:
//beer
//coffee
//tea
//water
//wine
//
//The french for beer is 啤酒
//The french for coffee is 咖啡
//The french for tea is 茶
//The french for water is 白开水
//The french for wine is 葡萄酒
