package main

import "fmt"

type obj interface{}

func main() {
	mf := func(i obj) obj {
		switch i.(type) {
		case int:
			return i.(int) * 2
		case string:
			return i.(string) + i.(string)
		}
		return i
	}

	isl := []obj{0, 1, 2, 3, 4, 5}
	res1 := mapFunc(mf, isl)
	for _, v := range res1 {
		fmt.Println(v)
	}

	ssl := []obj{"1", "2", "3", "4"}
	res2 := mapFunc(mf, ssl)
	for _, v := range res2 {
		fmt.Println(v)
	}

}

func mapFunc(mf func(obj) obj, list []obj) []obj {

	result := make([]obj, len(list))

	for ix, v := range list {
		result[ix] = mf(v)
	}

	return result
}

//0
//2
//4
//6
//8
//10
//11
//22
//33
//44
