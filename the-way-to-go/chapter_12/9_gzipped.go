package main

import (
	"bufio"
	"compress/gzip"
	"fmt"
	"os"
)

func main() {
	// 展示读取压缩文件的文本数据

	// 压缩包名
	fName := "products.txt.gz"

	// 定义一个读取器
	var r *bufio.Reader

	fi, err := os.Open(fName)

	if err != nil {
		_, err := fmt.Fprintf(os.Stderr, "%v, Can't open %s: error: %s\n", os.Args[0], fName, err)
		if err != nil {
			return
		}
	}
	defer func(fi *os.File) {
		err := fi.Close()
		if err != nil {
			return
		}
	}(fi)

	// 创建gzip的读取器
	fz, err := gzip.NewReader(fi)
	if err != nil {
		r = bufio.NewReader(fi)
	} else {
		r = bufio.NewReader(fz)
	}

	for {
		line, err := r.ReadString('\n')
		if err != nil {
			fmt.Println("Done reading file")
			os.Exit(0)
		}
		fmt.Println(line)
	}
}
