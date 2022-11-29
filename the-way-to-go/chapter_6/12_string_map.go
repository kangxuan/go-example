package main

import (
	"fmt"
	"strings"
)

func main() {
	asciiOnly := func(c rune) rune {
		if c > 127 {
			return ' '
		}
		return c
	}

	symbolConvert := func(c rune) rune {
		if c == '？' {
			return ' '
		}
		return c
	}

	fmt.Println(strings.Map(asciiOnly, "NTM的"))
	fmt.Println(strings.Map(symbolConvert, "你？"))
}
