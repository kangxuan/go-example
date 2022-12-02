package main

import (
	"bufio"
	"fmt"
	"os"
)

var inputReader *bufio.Reader

func main() {
	// 会创建一个读取器，并将其与标准输入绑定
	inputReader = bufio.NewReader(os.Stdin)

	fmt.Println("Please enter some input: ")

	input, err := inputReader.ReadString('\n')

	if err == nil {
		fmt.Println("The input was: ", input)
	}

}

//Please enter some input:
//shan la
//The input was:  shan la
