package main

import (
	"crypto/md5"
	"fmt"
	"io"
)

func main() {
	var b []byte
	hash := md5.New()

	_, err2 := io.WriteString(hash, "test")
	if err2 != nil {
		return
	}

	fmt.Printf("Result:%x\n", hash.Sum(b))
	fmt.Printf("Result:%d\n", hash.Sum(b))
}

//Result:098f6bcd4621d373cade4e832627b4f6
//Result:[9 143 107 205 70 33 211 115 202 222 78 131 38 39 180 246]
