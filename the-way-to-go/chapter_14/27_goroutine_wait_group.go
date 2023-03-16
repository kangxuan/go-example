package main

import (
	"fmt"
	"strconv"
	"sync"
)

// 当你并不关心并发操作的结果或者有其它方式收集并发操作的结果时，WaitGroup是实现等待一组并发操作完成的好方法。
var wg sync.WaitGroup

func helloWord() {
	fmt.Println("hello word!")
	// 告诉已经执行完成
	wg.Done()
}

func hello(i int) {
	fmt.Println("hello " + strconv.Itoa(i))
	defer wg.Done() // goroutine结束就登记-1
}

func main() {
	// 添加一个WaitGroup
	wg.Add(2)

	go helloWord()
	go helloWord()

	fmt.Println("你好，世界！")

	for i := 0; i < 4; i++ {
		wg.Add(1) // 启动一个goroutine就登记+1
		go hello(i)
	}
	// 等待所有登记的goroutine都结束
	wg.Wait()
}
