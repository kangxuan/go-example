package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	// 展示将一个文件的数据拷贝到另一个文件中
	inputFile := "product.txt"
	outputFile := "product_copy.txt"

	// 读取文件内容并存入到buf
	buf, err := ioutil.ReadFile(inputFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "File Error:%s\n", err)
	}
	fmt.Printf("%s\n", string(buf))
	// 将内容写入到文件中
	err = ioutil.WriteFile(outputFile, buf, 0644)
	if err != nil {
		panic(err.Error())
	}
}

//面包
//饺子
