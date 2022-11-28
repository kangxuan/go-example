package main

import (
	"fmt"
	"strings"
)

func main() {
	var str = "This is an example of a string"
	fmt.Printf("T/F? Does the string \"%s\" have prefix %s? ", str, "Th")
	// 是否以Th"开头
	fmt.Printf("%t\n", strings.HasPrefix(str, "Th"))
	fmt.Printf("T/F? Does the string \"%s\" have constains %s? ", str, "a string")
	// 是否包含a string"
	fmt.Printf("%t\n", strings.Contains(str, "a string"))
}
