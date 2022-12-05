package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// 展示读取命令行参数
	who := ""
	fmt.Printf("os.Args len: %d\n", len(os.Args))
	// os.Args如果没有参数，len(os.Args)==1,其中os.Args(0)是程序的名字
	if len(os.Args) > 1 {
		who += strings.Join(os.Args[1:], " ")
	}
	fmt.Println("Good Morning", who)
}

//go run 14_os_args.go Alice
//os.Args len: 2
//Good Morning Alice
