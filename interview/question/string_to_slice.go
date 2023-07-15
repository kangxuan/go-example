package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	s1 := "abcdefg"
	// 强制转换都会发生内存拷贝
	s2 := []byte(s1)
	fmt.Println(s2)
	// 规避发生内存拷贝
	// unsafe.Pointer(&s1)方法可以得到变量s1的地址
	// (*reflect.StringHeader)(unsafe.Pointer(&a))可以把字符串s1转成底层结构的形式
	ssh := *(*reflect.StringHeader)(unsafe.Pointer(&s1))
	// (*[]byte)(unsafe.Pointer(&ssh)) 可以把ssh底层结构体转成byte的切片的指针
	s3 := *(*[]byte)(unsafe.Pointer(&ssh))
	fmt.Printf("%v", s3)
}
