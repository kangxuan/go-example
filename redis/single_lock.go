package main

import (
	"fmt"
	"sync"
)

func main() {
	var lock sync.Mutex
	i := 0
	// 加锁
	lock.Lock()
	// 释放锁
	defer lock.Unlock()
	i++
	fmt.Println(i)
}
