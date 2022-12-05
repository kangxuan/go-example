package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// 展示如何向标准输入输出输出数据，

	// 1. 向标准输出直接输出数据
	_, err := fmt.Fprintf(os.Stdout, "%s\n", "hello world! - unbuffered")
	if err != nil {
		fmt.Println(err)
		return
	}

	// 2. 通过缓冲的形式向标准输出，输出数据
	buf := bufio.NewWriter(os.Stdout)
	_, err = fmt.Fprintf(buf, "%s", "hello world! - buffered")
	if err != nil {
		fmt.Println(err)
		return
	}

	// 将任何缓冲数据写入底层的io.Writer
	err = buf.Flush()
	if err != nil {
		fmt.Print(err)
		return
	}
}

//hello world! - unbuffered
//hello world! - buffered
