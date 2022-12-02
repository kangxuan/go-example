package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var nrchars, nrwords, nrlines = 0, 0, 0

func main() {
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter string:")

	for {
		input, err := inputReader.ReadString('\n')
		if err != nil {
			fmt.Printf("An error occurred: %s\n", err)
		}

		if input == "S\n" {
			fmt.Println("Here are the counts:")
			fmt.Printf("Number of characters: %d\n", nrchars)
			fmt.Printf("Number of words: %d\n", nrwords)
			fmt.Printf("Number of lines: %d\n", nrlines)
			os.Exit(0)
		}

		Counters(input)

	}

}

func Counters(input string) {
	nrchars += len(input) - 1 // 减去\n
	nrwords += len(strings.Fields(input))
	nrlines += 1
}

//Please enter string:
//fasdf
//adsf
//adfa
//S
//Here are the counts:
//Number of characters: 13
//Number of words: 3
//Number of lines: 3
