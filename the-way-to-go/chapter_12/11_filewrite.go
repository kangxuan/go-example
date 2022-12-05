package main

import (
	"os"
)

func main() {
	// 展示将数据输出到屏幕
	_, err := os.Stdout.WriteString("hello, world\n")
	if err != nil {
		return
	}

	outputFile, _ := os.OpenFile("test", os.O_CREATE|os.O_WRONLY, 0666)
	_, err2 := outputFile.WriteString("Hello, world\n")
	if err2 != nil {
		return
	}
	// 关闭文件
	defer func(outputFile *os.File) {
		err := outputFile.Close()
		if err != nil {
			return
		}
	}(outputFile)
}
