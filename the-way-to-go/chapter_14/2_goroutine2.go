package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)

	go setData(ch)
	go getData1(ch)

	time.Sleep(1 * time.Second)
}

func setData(ch chan string) {
	ch <- "Washington"
	ch <- "Tripoli"
	ch <- "London"
	ch <- "Beijing"
	ch <- "Tokyo"
}

func getData1(ch chan string) {
	var input string
	for {
		input = <-ch
		fmt.Printf("%s ", input)
	}
}

//Washington Tripoli London Beijing Tokyo
