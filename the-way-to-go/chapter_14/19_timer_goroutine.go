package main

import (
	"fmt"
	"time"
)

func main() {
	tick := time.Tick(time.Second)
	boom := time.After(5 * time.Second)

	for true {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("boom.")
			return
		default:
			fmt.Println("    .")
			time.Sleep(1e9)
		}
	}
}

//    .
//tick.
//    .
//tick.
//    .
//tick.
//    .
//tick.
//    .
//boom.
