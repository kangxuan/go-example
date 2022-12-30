package main

import (
	"fmt"
	"os"
)

// 展示如何读取文件内容
func main() {
	f := "test.txt"
	fl, err := os.Open(f)
	if err != nil {
		fmt.Println(f, err)
		return
	}
	defer fl.Close()
	buf := make([]byte, 1024)
	for {
		n, _ := fl.Read(buf)
		if 0 == n {
			break
		}
		os.Stdout.Write(buf[:n])
	}
}
