package main

import (
	"fmt"
	"os"
)

// tel1
func tel1(ch chan int, quit chan bool) {
	for i := 0; i < 15; i++ {
		ch <- i
	}
	quit <- true
}

func main() {
	ch := make(chan int)
	qu := make(chan bool)

	go tel1(ch, qu)

	for true {
		select {
		case i := <-ch:
			fmt.Printf("ok is %t and the counter is at %d\n", true, i)
		case <-qu:
			os.Exit(0)
		}
	}
}

//ok is true and the counter is at 0
//ok is true and the counter is at 1
//ok is true and the counter is at 2
//ok is true and the counter is at 3
//ok is true and the counter is at 4
//ok is true and the counter is at 5
//ok is true and the counter is at 6
//ok is true and the counter is at 7
//ok is true and the counter is at 8
//ok is true and the counter is at 9
//ok is true and the counter is at 10
//ok is true and the counter is at 11
//ok is true and the counter is at 12
//ok is true and the counter is at 13
//ok is true and the counter is at 14
