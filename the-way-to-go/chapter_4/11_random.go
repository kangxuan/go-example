package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	for i := 0; i < 10; i++ {
		a := rand.Int()
		fmt.Printf("%d / ", a)
	}
	fmt.Println()

	for i := 0; i < 5; i++ {
		r := rand.Intn(8)
		fmt.Printf("%d /", r)
	}
	fmt.Println()

	for i := 0; i < 5; i++ {
		b := rand.Float64()
		fmt.Printf("%f /", b)
	}
	fmt.Println()

	timens := int64(time.Now().Nanosecond())
	rand.Seed(timens)
	for i := 0; i < 10; i++ {
		fmt.Printf("%2.2f / ", 100*rand.Float32())
	}
}

// 5577006791947779410 / 8674665223082153551 / 6129484611666145821 / 4037200794235010051 / 3916589616287113937 / 6334824724549167320 / 605394647632969758 / 1443635317331776148 / 894385949183117216 / 2775422040480279449 /
// 6 /7 /2 /1 /0 /
// 0.468890 /0.283034 /0.293102 /0.679085 /0.218553 /
// 43.95 / 96.19 / 78.13 / 34.15 / 86.78 / 51.82 / 24.65 / 7.17 / 39.38 / 61.82 /
