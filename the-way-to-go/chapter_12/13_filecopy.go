package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// 展示将源文件的内容拷贝到目标文件中
	_, err := CopyFile("target.txt", "source.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Copy done!")
}

// CopyFile 复制源文件到目标文件
func CopyFile(dstFile, srcFile string) (written int64, err error) {
	// 打开源文件
	src, err := os.Open(srcFile)
	if err != nil {
		return
	}
	defer func(src *os.File) {
		err := src.Close()
		if err != nil {
			return
		}
	}(src)

	// 创建目标文件
	dst, err := os.Create(dstFile)
	if err != nil {
		return
	}
	defer func(dst *os.File) {
		err := dst.Close()
		if err != nil {
			return
		}
	}(dst)

	// 拷贝内容``
	return io.Copy(dst, src)
}

//Copy done!
