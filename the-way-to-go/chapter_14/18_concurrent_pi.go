package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	start := time.Now()
	fmt.Println(CalculatePi(5000))
	end := time.Now()
	delta := end.Sub(start)
	fmt.Printf("longCalculation took this amount of time: %s\n", delta)
}

// CalculatePi 计算π
func CalculatePi(n int) float64 {
	ch := make(chan float64)
	for k := 0; k <= n; k++ {
		go term(ch, float64(k))
	}

	f := 0.0
	for k := 0; k <= n; k++ {
		f += <-ch
	}
	return f
}

func term(ch chan float64, k float64) {
	ch <- 4 * math.Pow(-1, k) / (2*k + 1)
}

//3.1417926135957908
//longCalculation took this amount of time: 9.419302ms
