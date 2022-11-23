package main

import "fmt"

type TZ int

func main() {
	var a, b TZ = 3, 4
	c := a + b
	fmt.Printf("c has the value: %d\n", c)

	//	var f float32 = 10
	//	var f1 float64 = 10
	//	f1 = f
	//	fmt.Println(f1)
	//	var i int = 1
	//	var i1 int32 = 2
	//	i1 = i
	//	fmt.Println(i)

	e, f, g := 1, 2, 3
	fmt.Println(e, f, g)

	var b1 byte = 1
	fmt.Println(b1)

	var r1 rune = 12
	fmt.Println(r1)

	var s = "中国你好"
	// 先将s转成rune切片再求len
	sr := []rune(s)
	fmt.Println(len(sr))
	fmt.Println(string(sr[0:1]))

	var f1 = -1.0000000000001000
	fmt.Printf("f1 : %g, %G, %f\n", f1, f1, f1)

	//var n int = 3000
	var n float32 = 30.00
	n1 := int(n)
	// 补0
	fmt.Printf("n :%05d\n", n1)
}

// c has the value: 7
//1 2 3
//1
//12
//4
//中
//f1 : -1.0000000000001, -1.0000000000001, -1.000000
//n :00030
