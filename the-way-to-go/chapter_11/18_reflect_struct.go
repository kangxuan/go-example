package main

import (
	"fmt"
	"reflect"
)

type NotknownType struct {
	s1, s2, S3 string
}

// String 如果类型有String()方法，通过fmt.Print打印时会调用这个方法
func (n NotknownType) String() string {
	return n.s1 + "-" + n.s2 + "-" + n.S3
}

var secret interface{} = NotknownType{"Ada", "Go", "Oberon"}

func main() {
	value := reflect.ValueOf(secret)
	typ := reflect.TypeOf(secret)
	fmt.Println(value)
	fmt.Println(typ)

	knd := value.Kind() // struct
	fmt.Println(knd)

	// v.NumField()返回字段数量
	// v.Field(i)返回具体某一个字段的值
	for i := 0; i < value.NumField(); i++ {
		fmt.Printf("Field %d: %v\n", i, value.Field(i))
	}

	// Method(i)返回具体某一个方法，Call()调用此方法
	results := value.Method(0).Call(nil)
	fmt.Println(results) // [Ada - Go - Oberon]

	// 因为Filed(0)不是导出字段，所以通过SetString不允许设置
	//panic: reflect: reflect.Value.SetString using value obtained using unexported field
}

//Ada-Go-Oberon
//main.NotknownType
//struct
//Field 0: Ada
//Field 1: Go
//Field 2: Oberon
//[Ada-Go-Oberon]
