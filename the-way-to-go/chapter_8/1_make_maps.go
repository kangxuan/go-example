package main

import "fmt"

func main() {
	// 定义一个key为string，值为int的map
	var mapLit map[string]int
	var mapAssigned map[string]int

	// 为mapLit赋值
	mapLit = map[string]int{"one": 1, "two": 2}
	// 通过make方式定义
	mapCreated := make(map[string]float32)
	// mapLit赋值给mapAssigned
	mapAssigned = mapLit

	mapCreated["key1"] = 11
	mapCreated["key2"] = 22.22
	mapAssigned["two"] = 3

	fmt.Printf("Map literal at \"one\" is: %d\n", mapLit["one"])
	fmt.Printf("Map literal at \"two\" is: %d\n", mapLit["two"]) // 注意此处输出3，因为map和切片一样都是索引类型，共享了内存，修改了mapAssigned的值同时mapLit也会受到影响
	fmt.Printf("Map created at \"key1\" is: %f\n", mapCreated["key1"])
	fmt.Printf("Map created at \"key2\" is: %f\n", mapCreated["key2"])
	fmt.Printf("Map assigned at \"two\" is: %d\n", mapLit["two"])
	fmt.Printf("Map literal at \"ten\" is: %d\n", mapLit["ten"])
	//	fmt.Printf("Map literal at \"1\" is: %d\n", mapLit[1]) // 如果key为string，那么访问用int会报错

	if _, ok := mapLit["ten"]; ok {
		fmt.Println(mapLit["ten"])
	} else {
		fmt.Printf("Map literal at \"ten\" is not existed\n")
	}
}

//Map literal at "one" is: 1
//Map literal at "two" is: 3
//Map created at "key1" is: 11.000000
//Map created at "key2" is: 22.219999
//Map assigned at "two" is: 3
//Map literal at "ten" is: 0
//Map literal at "ten" is not existed
