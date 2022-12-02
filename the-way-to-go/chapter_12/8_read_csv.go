package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

// 定义一个Book结构体
type Book struct {
	title    string  // 书名
	price    float64 // 价格
	quantity int     // 数量
}

func main() {
	// 定义一个Book切片
	var bks = make([]Book, 1)

	// 打开文件
	file, err := os.Open("products.txt")
	if err != nil {
		log.Fatalf("Error %s opening file products.txt: ", err)
	}
	// 程序关闭前，关闭句柄
	defer file.Close()

	// 创建一个阅读器
	reader := bufio.NewReader(file)
	// 循环读出每一行进行处理
	for {
		line, err := reader.ReadString('\n')
		readErr := err
		if readErr == io.EOF {
			break
		}

		// 去掉\n
		line = line[:len(line)-1]
		// 通过;分割
		strSl := strings.Split(line, ";")

		// 实例化一个Book
		book := new(Book)
		// 标题
		book.title = strSl[0]
		// 价格并转化成对应类型
		book.price, err = strconv.ParseFloat(strSl[1], 64)
		if err != nil {
			fmt.Printf("Error in file: %v\n", err)
		}
		// 数量并转化成对应类型
		book.quantity, err = strconv.Atoi(strSl[2])
		if err != nil {
			fmt.Printf("Error in file: %v\n", err)
		}

		if bks[0].title == "" {
			bks[0] = *book
		} else {
			bks = append(bks, *book)
		}
	}

	fmt.Println("We have read the following books from the file: ")
	for _, bk := range bks {
		fmt.Println(bk)
	}
}

//We have read the following books from the file:
//{"The ABC of Go" 25.5 1500}
//{"Functional Programming with Go" 56 280}
//{"Go for It" 45.9 356}
//{"The Go Way" 55 500}
