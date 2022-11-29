package main

import "fmt"

func main() {
	str := "1aabc"
	s := []byte(str)
	fmt.Println(s)

	result := getDiffArray(s)
	fmt.Println(string(result))
}

func getDiffArray(s []byte) []byte {
	result := make([]byte, len(s))
	j := 0
	for i := 1; i < len(s); i++ {
		if s[i-1] != s[i] {
			result[j] = s[i]
			j++
		}
	}

	return result
}
