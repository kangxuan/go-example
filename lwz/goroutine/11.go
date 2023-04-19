package main

import (
	"fmt"
	"sync"
)

var (
	x1  int
	wg1 sync.WaitGroup
	m1  sync.Mutex
)

func addSync() {
	for i := 0; i < 5000; i++ {
		// 修改x前加锁
		m1.Lock()
		x1 = x1 + i
		// 改完解锁
		m1.Unlock()
	}
	wg1.Done()
}

func main() {
	wg1.Add(2)

	go addSync()
	go addSync()

	wg1.Wait()
	fmt.Println(x1)
}
