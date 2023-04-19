package main

import (
	"fmt"
	"sync"
)

var (
	x  int
	wg sync.WaitGroup
)

func add() {
	for i := 0; i < 5000; i++ {
		x = x + i
	}
	wg.Done()
}

func main() {
	wg.Add(2)

	go add()
	go add()

	wg.Wait()
	fmt.Println(x)
}
