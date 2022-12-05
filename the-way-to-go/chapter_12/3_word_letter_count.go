package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var nrchars, nrwords, nrlines = 0, 0, 0

func main() {
	// 创建一个读取器，并将其与标准输入绑定
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter string:")

	// 循环监听标准输入
	for {
		input, err := inputReader.ReadString('\n')
		if err != nil {
			fmt.Printf("An error occurred: %s\n", err)
		}

		// 当以"S"结尾，则输出统计数据
		if input == "S\n" {
			fmt.Println("Here are the counts:")
			fmt.Printf("Number of characters: %d\n", nrchars)
			fmt.Printf("Number of words: %d\n", nrwords)
			fmt.Printf("Number of lines: %d\n", nrlines)
			os.Exit(0)
		}

		// 统计数量
		Counters(input)
	}

}

// Counters 统计字符串的数量
func Counters(input string) {
	nrchars += len(input) - 1             // 字符个数
	nrwords += len(strings.Fields(input)) // 单词数
	nrlines += 1                          // 行数
}

//Please enter string:
//这是第一行
//这是第二行
//这是第三行
//S
//Here are the counts:
//Number of characters: 45
//Number of words: 3
//Number of lines: 3
