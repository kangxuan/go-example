package main

import "fmt"

func main() {
	//var map1 map[int]float32
	map1 := make(map[int]float32)
	map1[1] = 1.1
	map1[2] = 2.2
	map1[3] = 3.3
	map1[4] = 4.4

	for key, value := range map1 {
		fmt.Printf("map1 key:%d, value:%f\n", key, value)
	}
}

// map 是无序的
//map1 key:3, value:3.300000
//map1 key:4, value:4.400000
//map1 key:1, value:1.100000
//map1 key:2, value:2.200000
