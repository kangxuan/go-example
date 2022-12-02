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

	res1 := mapFunc(mf, 1, 2, 3, 4, 5)
	for _, v := range res1 {
		fmt.Println(v)
	}

	res2 := mapFunc(mf, "a", "b", "c")
	for _, v := range res2 {
		fmt.Println(v)
	}

}

func mapFunc(mf func(obj) obj, list ...obj) []obj {

	result := make([]obj, len(list))

	for ix, v := range list {
		result[ix] = mf(v)
	}

	return result
}
