package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"fmt"
	"io"
)

func main() {
	// 基础方案 - 单向哈希
	// sha256
	h := sha256.New()
	io.WriteString(h, "sdfj12323")
	fmt.Printf("%x \n", h.Sum(nil))

	// sha1
	h = sha1.New()
	io.WriteString(h, "sdfj12323")
	fmt.Printf("%x \n", h.Sum(nil))

	// md5
	h = md5.New()
	io.WriteString(h, "sdfj12323")
	fmt.Printf("%x \n", h.Sum(nil))

	// 进阶方案 - 加盐
	h = md5.New()
	io.WriteString(h, "sdfj12323")
	fmt.Printf("%x \n", h.Sum(nil))

	// 加盐
	io.WriteString(h, "sd1d")
	fmt.Printf("%x \n", h.Sum(nil))

	// 继续加盐
	io.WriteString(h, "sd11")
	fmt.Printf("%x \n", h.Sum(nil))

	// 专家方案 -
}
