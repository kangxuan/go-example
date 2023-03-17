package main

import "fmt"

func main() {
	ch := make(chan int, 2)
	for i := 1; i <= 10; i++ {
		select {
		case x := <-ch:
			fmt.Print(x, " ")
		case ch <- i:
		}
	}
}

// i=1 ch<-1 1
// i=2 1, ch<-2 2
// i=3 2, ch<-3 3
