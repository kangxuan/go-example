package main

import "fmt"

func main() {
	//LABAL1:
	//for {
	//	for i:=0; i<10; i++ {
	//		if(i==3) {
	//			break LABAL1
	//		}
	//	}
	//}
	//fmt.Println("OK")

LABAL1:
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		for {
			continue LABAL1
		}
	}
	fmt.Println("ok")
}
