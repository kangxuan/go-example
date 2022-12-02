package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	// 打开文件
	inputFile, inputError := os.Open("input.dat")

	if inputError != nil {
		fmt.Printf("An error occurred on opening the inputfile\n" +
			"Does the file exist?\n" +
			"Have you got access to it?\n")
		return
	}

	defer inputFile.Close() // 确保在程序退出前关闭文件
	// 创建读取器，并将其与文件句柄绑定
	inputReader := bufio.NewReader(inputFile)
	for {
		// 一行一行读
		inputString, readerError := inputReader.ReadString('\n')
		//inputString, _, readerError := inputReader.ReadLine()
		fmt.Printf("The input was : %s", inputString)

		if readerError == io.EOF {
			return
		}
	}

}

//The input was : hello
//The input was : world!
//The input was : 真他妈服了
//The input was : %
