package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	//var num float64 = 0.25

	fmt.Print("First example with -1: ")
	ret1, err1 := MySqrt(-1)
	if err1 != nil {
		fmt.Println("Error! Return values are: ", ret1, err1)
	} else {
		fmt.Println("It's ok! Return values are: ", ret1, err1)
	}

	fmt.Print("Second example with 5: ")
	//you could also write it like this
	if ret2, err2 := MySqrt(5); err2 != nil {
		fmt.Println("Error! Return values are: ", ret2, err2)
	} else {
		fmt.Println("It's ok! Return values are: ", ret2, err2)
	}

	fmt.Println(MySqrt_2(0.25))
}

func MySqrt(num float64) (float64, error) {
	if num < 0 {
		return float64(math.NaN()), errors.New("I won't be able to do a sqrt of negative number!")
	}

	return math.Sqrt(num), nil
}

func MySqrt_2(num float64) (sqrt float64, err error) {

	if num < 0 {
		sqrt = float64(math.NaN())
		err = errors.New("I won't be able to do a sqrt of negative number!")
	}

	sqrt = math.Sqrt(num)
	err = nil

	return

}
