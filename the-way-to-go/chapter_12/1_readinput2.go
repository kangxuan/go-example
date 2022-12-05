package main

import (
	"bufio"
	"fmt"
	"os"
)

// 声明一个缓冲的读取器
var inputReader *bufio.Reader

func main() {
	// 创建一个读取器，并将其与标准输入绑定
	inputReader = bufio.NewReader(os.Stdin)
	fmt.Println("Please enter some input: ")

	// 从标准输入中读取文本并以换行结束
	input, err := inputReader.ReadString('\n')
	if err == nil {
		fmt.Println("The input was: ", input)
	}
}

//Please enter some input:
//shan la
//The input was:  shan la
