package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println(log.Flags())
	log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
	log.Println("这是一条很普通的日志。")
}
