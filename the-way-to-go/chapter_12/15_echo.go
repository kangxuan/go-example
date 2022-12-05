package main

import (
	"flag"
	"os"
)

// NewLine 定义参数-n==false
var NewLine = flag.Bool("n", false, "print newline")

const (
	Space   = " "
	Newline = "\n"
)

func main() {
	// 展示读取命令行参数
	// 打印帮助信息
	flag.PrintDefaults()
	// 扫描参数列表并设置标志
	flag.Parse()

	var s string

	// flag.NArg()返回参数的个数
	for i := 0; i < flag.NArg(); i++ {
		if i > 0 {
			s += Space
			// 如果有传-n
			if *NewLine {
				s += Newline
			}
		}
		// flag.Arg(i)
		s += flag.Arg(i)
	}

	// 输出到屏幕
	_, err2 := os.Stdout.WriteString(s)
	if err2 != nil {
		return
	}
}
