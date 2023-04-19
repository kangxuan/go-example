package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	x2       int64
	wg2      sync.WaitGroup
	mutex2   sync.Mutex
	rwMutex2 sync.RWMutex
)

func writeWithLock() {
	// 加互斥锁
	mutex2.Lock()
	x2 += 1
	time.Sleep(10 * time.Millisecond) //假设写操作耗时10毫秒
	// 解开互斥锁
	mutex2.Unlock()
	wg2.Done()
}

func readWithLock() {
	// 加互斥锁
	mutex2.Lock()
	time.Sleep(10 * time.Millisecond) //假设读操作耗时10毫秒
	// 解开互斥锁
	mutex2.Unlock()
	wg2.Done()
}

func writeWithRWLock() {
	// 加写锁
	rwMutex2.Lock()
	x2 += 1
	time.Sleep(10 * time.Millisecond) //假设写操作耗时10毫秒
	// 解写锁
	rwMutex2.Unlock()
	wg2.Done()
}

func readWithRWLock() {
	// 加读锁
	rwMutex2.RLock()
	time.Sleep(10 * time.Millisecond) //假设读操作耗时10毫秒
	// 解读锁
	rwMutex2.RUnlock()
	wg2.Done()
}

func do(wf, rf func(), wTime, rTime int) {
	start := time.Now()
	// 并发wTime次wf
	for i := 0; i < wTime; i++ {
		wg2.Add(1)
		go wf()
	}

	// 并发rTime次rf
	for i := 0; i < rTime; i++ {
		wg2.Add(1)
		go rf()
	}

	wg2.Wait()
	cost := time.Since(start)
	fmt.Printf("x:%v cost:%v\n", x2, cost)
}

func main() {
	do(writeWithLock, readWithLock, 10, 1000)

	do(writeWithRWLock, readWithRWLock, 10, 1000)
}
