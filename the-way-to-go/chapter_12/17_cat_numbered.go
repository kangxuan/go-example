package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

var numberFlag = flag.Bool("n", false, "number each line")

// cat1 读取数据
func cat1(r *bufio.Reader) {
	i := 1
	for {
		// 一行一行读取
		buf, err := r.ReadBytes('\n')
		if err == io.EOF {
			break
		}

		if *numberFlag {
			_, err = fmt.Fprintf(os.Stdout, "%5d %s", i, buf)
			i++
		} else {
			_, err = fmt.Fprintf(os.Stdout, "%s", buf)
		}
	}
}

func main() {
	// 扫描命令行册数
	flag.Parse()
	if flag.NArg() == 0 {
		cat1(bufio.NewReader(os.Stdin))
	}

	for i := 0; i < flag.NArg(); i++ {
		f, err := os.Open(flag.Arg(i))
		if err != nil {
			_, err = fmt.Fprintf(os.Stderr, "%s:error reading from %s: %s\n", os.Args[0], flag.Arg(i), err.Error())
			continue
		}
		cat1(bufio.NewReader(f))

		err = f.Close()
		if err != nil {
			return
		}
	}
}
