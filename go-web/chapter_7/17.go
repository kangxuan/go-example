package main

import (
	"fmt"
	"os"
)

// 展示目录的操作
func main() {
	// 创建目录
	os.Mkdir("dir1", 0777)
	// 递归创建目录
	os.MkdirAll("dir1/dir2/dir3", 0777)

	// 删除目录
	err := os.Remove("dir1")
	if err != nil {
		fmt.Println(err)
	}
	// 递归删除目录
	os.RemoveAll("dir1")
}
