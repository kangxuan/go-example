package main

import (
	"fmt"
	"math"
	"math/big"
)

func main() {
	im := big.NewInt(math.MaxInt64)
	in := im

	io := big.NewInt(1956)
	ip := big.NewInt(1)
	ip.Mul(im, in).Add(ip, im).Div(ip, io)
	fmt.Printf("Big Int: %v\n", ip)
}
