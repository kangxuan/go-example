package main

import (
	"./struct_pack"
	"fmt"
)

func main() {
	struct1 := new(struct_pack.ExpStruct)
	struct1.Mi1 = 10
	struct1.Mf1 = 16.
	fmt.Printf("Mi1 = %d\n", struct1.Mi1)
	fmt.Printf("Mf1 = %f\n", struct1.Mf1)
}

//Mi1 = 10
//Mf1 = 16.000000
