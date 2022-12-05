package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// 展示利用缓冲区，写入数据到文件
	// 打开或创建写入的文件
	outputFile, outputErr := os.OpenFile("output.dat", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if outputErr != nil {
		fmt.Println("An error occurred with file opening or creation")
		return
	}
	defer func(outputFile *os.File) {
		err := outputFile.Close()
		if err != nil {
			return
		}
	}(outputFile)

	// 新建一个写入器（缓冲区）对象并与outputFile绑定
	outputWriter := bufio.NewWriter(outputFile)
	outputString := "hello world，Thank!\n"

	for i := 0; i < 10; i++ {
		// 写入数据到缓冲区
		_, err2 := outputWriter.WriteString(outputString)
		if err2 != nil {
			return
		}
	}

	// 缓冲区的内容紧接着被完全写入文件
	err := outputWriter.Flush()
	if err != nil {
		return
	}
}
