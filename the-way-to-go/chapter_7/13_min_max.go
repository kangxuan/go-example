package main

import "fmt"

func main() {
	var a = []int{1, 3, 2, 13, 5, 24235}

	fmt.Printf("a'max:%d", getMaxOrMin(a, 1))
}

func getMaxOrMin(a []int, getType int) int {
	var temp int
	for _, value := range a {
		if getType == 1 {
			if temp < value {
				temp = value
			}
		} else {
			if temp > value {
				temp = value
			}
		}
	}

	return temp
}
