package main

import (
	"float"
	"fmt"
)

func main() {
	f1 := float.NewFloat64Array()
	f1.Fill(10)

	fmt.Printf("Before sorting %s\n", f1)
	float.Sort(f1)
	fmt.Printf("After sorting %s\n", f1)

	if float.IsSorted(f1) {
		fmt.Println("The float64 array is sorted!")
	} else {
		fmt.Println("The float64 array is NOT sorted!")
	}
}

//Before sorting { 75.8 46.2 25.5 31.9 8.3 1.7 30.3 41.7 22.3 96.6  }
//After sorting { 1.7 8.3 22.3 25.5 30.3 31.9 41.7 46.2 75.8 96.6  }
//The float64 array is sorted!
