package main

import "fmt"
import "strings"

func main() {
	var str = "My Name is Shanla, Shanla is my Name!"

	replace_str := strings.Replace(str, "Shanla", "Xuan", -1)

	fmt.Printf("The replace string is: %s\n", replace_str)
}
