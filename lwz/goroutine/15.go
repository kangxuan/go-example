package main

import (
	"fmt"
	"strconv"
	"sync"
)

// 并发安全的map
var syncM = sync.Map{}

func main() {
	wg := sync.WaitGroup{}
	// 对syncM执行20个并发的读写操作
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(n int) {
			key := strconv.Itoa(n)
			syncM.Store(key, n)         // 存储key-value
			value, _ := syncM.Load(key) // 根据key取值
			fmt.Printf("k=:%v,v:=%v\n", key, value)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
