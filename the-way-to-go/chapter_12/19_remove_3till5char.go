package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	// 展示将源文件数据写入到目标文件中
	// 打开文件
	inputFile, _ := os.Open("goprogram")
	// 创建并打开输出文件
	outputFile, _ := os.OpenFile("goprogramT", os.O_WRONLY|os.O_CREATE, 0666)
	// 程序退出前关闭文件
	defer func(inputFile *os.File) {
		err := inputFile.Close()
		if err != nil {
			return
		}
	}(inputFile)
	defer func(outputFile *os.File) {
		err := outputFile.Close()
		if err != nil {
			return
		}
	}(outputFile)

	// 创建一个inputFile的阅读器
	inputReader := bufio.NewReader(inputFile)
	// 创建一个outputFile的写入器
	outputWriter := bufio.NewWriter(outputFile)

	for {
		// 循环读取每行内容
		inputString, _, readerError := inputReader.ReadLine()
		if readerError == io.EOF {
			fmt.Println("EOF")
			break
		}

		// 将inputString的部分数据写入到
		outputString := string([]byte(inputString)[2:5]) + "\r\n"
		fmt.Println(outputString)
		_, err := outputWriter.WriteString(outputString)

		if err != nil {
			fmt.Println(err)
			return
		}
	}

	// 将缓冲区数据刷入文件
	err = outputWriter.Flush()
	if err != nil {
		fmt.Println("写入文件失败")
	}

	fmt.Println("Conversion done")
}
