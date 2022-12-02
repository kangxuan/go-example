package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x float64 = 11
	fmt.Println("type: ", reflect.TypeOf(x))

	v := reflect.ValueOf(x)
	fmt.Println("value: ", v)

	fmt.Println("type: ", v.Type())
	fmt.Println("kind: ", v.Kind())
	fmt.Println("value: ", v.Float())
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
