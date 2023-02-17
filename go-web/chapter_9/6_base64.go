package main

import (
	"encoding/base64"
	"fmt"
)

// base64Encode base64加密
func base64Encode(src []byte) []byte {
	return []byte(base64.StdEncoding.EncodeToString(src))
}

// base64Decode base64解密
func base64Decode(src []byte) ([]byte, error) {
	return base64.StdEncoding.DecodeString(string(src))
}

func main() {
	hello := "Hello world! 你好，世界！"
	debyte := base64Encode([]byte(hello))
	fmt.Printf("%s 加密后：%x\n", hello, debyte)

	enByte, _ := base64Decode(debyte)
	fmt.Printf("%x 解密后：%s", debyte, string(enByte))
}
