package main

import "fmt"

func generate(ch chan int) {
	for i := 2; ; i++ {
		ch <- i
	}
}

func filter(in, out chan int, prime int) {
	for true {
		i := <-in
		if i%prime != 0 {
			out <- i
		}
	}
}

func main() {
	// 展示打印所有素数
	ch := make(chan int)
	go generate(ch)

	for {
		prime := <-ch
		fmt.Print(prime, " ")
		ch1 := make(chan int)
		go filter(ch, ch1, prime)
		ch = ch1
	}
}

// 2
// 2
// 3 2
// out = 3
//
