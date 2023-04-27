package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	file, err := os.Open("./file.go")
	if err != nil {
		log.Fatalf("os.Open Error:%v\n", err)
	}
	defer file.Close()

	var temp = make([]byte, 500)
	n, err := file.Read(temp)
	if err == io.EOF {
		log.Println("文件已读完")
		return
	}
	if err != nil {
		log.Fatalln("file.Read, err:", err)
	}
	fmt.Printf("读取了%d字节数据\n", n)
	fmt.Println(string(temp[:n]))
}
