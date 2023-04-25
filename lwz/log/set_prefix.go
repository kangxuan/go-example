package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println(log.Prefix())

	log.SetPrefix("[info]")
	log.Println("这是一条普通日志")
	log.SetPrefix("[error]")
	log.Println("这是一条错误日志")
}
