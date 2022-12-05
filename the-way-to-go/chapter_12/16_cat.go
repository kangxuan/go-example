package main

import (
	"flag"
	"fmt"
	"os"
)

const NBUF = 512

// cat 读取数据
func cat(f *os.File) {
	var buf [NBUF]byte
	for {
		// 每次读512个byte
		switch nr, err := f.Read(buf[:]); true {
		case nr < 0: // 如果数据错误则报错
			_, err = fmt.Fprintf(os.Stderr, "cat: error reading: %s\n", err.Error())
			os.Exit(1)
		case nr == 0: // 如果没有读取到数据则结束
			return
		case nr > 0: // 有数据则写入到
			if nw, ew := os.Stdout.Write(buf[0:nr]); nw != nr {
				_, err = fmt.Fprintf(os.Stderr, "cat: error writing: %s\n", ew.Error())
			}
		}
	}
}

func main() {
	// 展示通过输入文件名读取文件内容
	// 扫描命令行册数
	flag.Parse()
	if flag.NArg() == 0 {
		cat(os.Stdin)
	}

	// 循环参数
	for i := 0; i < flag.NArg(); i++ {
		f, err := os.Open(flag.Arg(i))
		if f == nil {
			_, err = fmt.Fprintf(os.Stderr, "cat: can't open %s: error %s\n", flag.Arg(i), err)
			os.Exit(1)
		}
		// 读取文件数据
		cat(f)
		// 关闭文件
		err = f.Close()
		if err != nil {
			return
		}
	}
}

//go run 16_cat.go product.txt
//面包
//饺子
