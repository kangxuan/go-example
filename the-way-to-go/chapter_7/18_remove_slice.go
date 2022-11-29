package main

import (
	"errors"
	"fmt"
)

func main() {
	sl := []string{"a", "b", "c", "d", "e"}

	result, err := RemoveStringSlice(sl, 1, 3)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

}

func RemoveStringSlice(sl []string, start int, end int) (result []string, err error) {
	if start >= end || start < 0 || end > len(sl) {
		err = errors.New("start或end错误")
	} else {
		result = make([]string, len(sl)-(end-start))
		at := copy(result, sl[:start])
		copy(result[at:], sl[end:])
	}

	return
}
