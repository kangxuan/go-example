package main

import (
	"bufio"
	"encoding/gob"
	"fmt"
	"log"
	"os"
)

type Address1 struct {
	Type    string
	City    string
	Country string
}

type VCard1 struct {
	FirstName string
	LastName  string
	Addresses []Address1
	Remark    string
}

var vc VCard1

func main() {
	// 展示gob编码和解码
	// 打开文件
	file, _ := os.Open("vcard.gob")
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			return
		}
	}(file)

	// 创建一个阅读器
	inReader := bufio.NewReader(file)
	// gob解码器
	dec := gob.NewDecoder(inReader)
	// 解码
	err := dec.Decode(&vc)
	if err != nil {
		log.Println("Error in decoding gob")
	}

	fmt.Println(vc)
}
