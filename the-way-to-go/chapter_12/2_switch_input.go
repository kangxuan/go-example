package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// 会创建一个读取器，并将其与标准输入绑定
	inputReader := bufio.NewReader(os.Stdin)

	fmt.Println("Please enter some input: ")

	// 方法 ReadString(delim byte)，该方法从输入中读取内容，直到碰到 delim 指定的字符，然后将读取到的内容连同 delim 字符一起放到缓冲区。
	input, err := inputReader.ReadString('\n')

	if err != nil {
		fmt.Println("There were errors reading, exiting program.")
		return
	}

	fmt.Printf("Your name is %s", input)
	// 类Unix的换行符是\n,windows的换行符是\r\n
	switch input {
	case "Philip\n":
		fmt.Println("Welcome Philip!")
	case "Chris\n":
		fmt.Println("Welcome Chris!")
	case "Ivo\n":
		fmt.Println("Welcome Ivo!")
	default:
		fmt.Println("You are not welcome here! Goodbye!")
	}
	// switch 2
	switch input {
	case "Philip\n":
		fallthrough
	case "Chris\n":
		fallthrough
	case "Ivo\n":
		fmt.Printf("Welcome %s\n", input)
	default:
		fmt.Println("You are not welcome here! Goodbye!")
	}
	// switch 3 - 针对windows系统\r\n
	switch input {
	case "Philip\r\n", "Ivo\r\n":
		fmt.Printf("Welcome %s\n", input)
	default:
		fmt.Println("You are not welcome here! Goodbye!")
	}
}
