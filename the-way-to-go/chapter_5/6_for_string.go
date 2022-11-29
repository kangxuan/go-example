package main

import "fmt"

func main() {
	var str1 = "1 324324 45253"
	for i := 0; i < len(str1); i++ {
		fmt.Printf("Character on position %d is: %c \n", i, str1[i])
	}

	str2 := "日本語"
	fmt.Printf("The length of str2 is: %d\n", len(str2))
	for ix := 0; ix < len(str2); ix++ {
		fmt.Printf("Character on position %d is: %c \n", ix, str2[ix])
	}
}
