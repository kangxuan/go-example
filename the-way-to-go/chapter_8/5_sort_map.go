package main

import (
	"fmt"
	"sort"
)

// 定义一个map
var barVal = map[string]int{
	"alpha": 34, "bravo": 56, "charlie": 23,
	"delta": 87, "echo": 56, "foxtrot": 12,
	"golf": 34, "hotel": 16, "indio": 87,
	"juliet": 65, "kili": 43, "lima": 98,
}

func main() {
	fmt.Println("排序前：")
	for k, v := range barVal {
		fmt.Printf("Key: %v, Value:%v / ", k, v)
	}

	// 定义一个切片用来放barVal的key
	keys := make([]string, len(barVal))
	i := 0
	for key := range barVal {
		keys[i] = key
		i++
	}

	// 对key的切片进行排序
	sort.Strings(keys)
	fmt.Println()
	fmt.Println("排序后：")
	for _, k := range keys {
		fmt.Printf("Key: %v, Value:%v / ", k, barVal[k])
	}
}

//排序前：
//Key: lima, Value:98 / Key: alpha, Value:34 / Key: bravo, Value:56 / Key: delta, Value:87 / Key: golf, Value:34 / Key: hotel, Value:16 / Key: juliet, Value:65 / Key: kili, Value:43 / Key: charlie, Value:23 / Key: echo, Value:56 / Key: foxtrot, Value:12 / Key: indio, Value:87 /
//排序后：
//Key: alpha, Value:34 / Key: bravo, Value:56 / Key: charlie, Value:23 / Key: delta, Value:87 / Key: echo, Value:56 / Key: foxtrot, Value:12 / Key: golf, Value:34 / Key: hotel, Value:16 / Key: indio, Value:87 / Key: juliet, Value:65 / Key: kili, Value:43 / Key: lima, Value:98 /
