package main

import "fmt"

func main() {
	m1 := map[int]string{0: "a", 1: "b", 2: "c"}
	m2 := map[string]int{}
	for k, v := range m1 {
		m2[v] = k
	}
	fmt.Println(m1)
	fmt.Println(m2)
}

//map[0:a 1:b 2:c]
//map[a:0 b:1 c:2]
