package main

import (
	"./pack1"
	"fmt"
)

func main() {
	var test1 string
	test1 = pack1.ReturnStr()
	fmt.Printf("ReturnStr form packge1:%s\n", test1)
	fmt.Printf("Integer from package1:%d\n", pack1.Pack1Int)
}
