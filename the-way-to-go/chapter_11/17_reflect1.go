package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x float64 = 11
	// 反射类型
	t := reflect.TypeOf(x)
	fmt.Println("type: ", t)
	fmt.Println("Name: ", t.Name())
	fmt.Println("String: ", t.String())
	//fmt.Println("...", t.Key())

	// 反射值
	v := reflect.ValueOf(x)
	fmt.Println("value: ", v)
	// 返回对应的v的类型
	fmt.Println("type: ", v.Type())
	// 返回对应v的最底层类型
	fmt.Println("kind: ", v.Kind())
	// 如果是float类型，通过Float()获取值
	fmt.Println("value: ", v.Float())
	// Interface将v的当前值作为一个interface{}返回
	fmt.Println("interface: ", v.Interface())
	fmt.Printf("value is %5.2e\n", v.Interface())
	y := v.Interface().(float64)
	fmt.Println(y)
}

//type:  float64
//value:  11
//type:  float64
//kind:  float64
//value:  11
//interface:  11
//value is 1.10e+01
//11
