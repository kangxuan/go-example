package main

import (
	"fmt"
	"os"
)

func main() {
	str := "hello 沙河"
	err := os.WriteFile("./xx.txt", []byte(str), 0666)
	if err != nil {
		fmt.Println("write file failed, err:", err)
		return
	}
}
