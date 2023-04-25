package main

import (
	"log"
	"os"
)

func main() {
	logger := log.New(os.Stdout, "info", log.Lshortfile|log.Ldate|log.Ltime)
	logger.Println("这是自定义的logger记录的日志。")
}
