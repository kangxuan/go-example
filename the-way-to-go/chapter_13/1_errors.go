package main

import (
	"errors"
	"fmt"
)

var errNotFound = errors.New("not found error")

func main() {
	// 展示如何生成一个errors对象
	fmt.Printf("error: %v", errNotFound)
}

//error: not found error
