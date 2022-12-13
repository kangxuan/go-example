package main

import (
	"fmt"
	"testing"
)

func main() {
	fmt.Println(" sync", testing.Benchmark(BenchmarkChannelSync).String())
	fmt.Println("buffered", testing.Benchmark(BenchmarkChannelBuffered).String())
}

func BenchmarkChannelSync(b *testing.B) {
	ch := make(chan int)
	go func() {
		for i := 0; i < b.N; i++ {
			ch <- i
		}
		close(ch)
	}()

	for range ch {

	}
}

func BenchmarkChannelBuffered(b *testing.B) {
	ch := make(chan int, 120)
	go func() {
		for i := 0; i < b.N; i++ {
			ch <- i
		}
		close(ch)
	}()

	for range ch {

	}
}

// sync  5199277         221.8 ns/op
//buffered 19974562               60.19 ns/op
