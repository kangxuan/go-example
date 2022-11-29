package main

import "fmt"

func main() {
	// 数组
	//	var a = [...]float32{1.0, 1.3, 2.4, 3.4}
	// 切片
	var a = []float32{1.0, 1.3, 2.4, 3.4}

	fmt.Printf("a's sum:%f\n", Sum(a))

	var b = []int{1, 2, 3}
	sum, average := SumAndAverage(b)
	fmt.Printf("b's sum:%d, b's average:%f", sum, average)
}

/*func Sum(a [4]float32) (sum float32) {
	for _, item := range a {
		sum += item
	}
	return
}*/
// 求和
func Sum(a []float32) (sum float32) {
	for _, value := range a {
		sum += value
	}

	return
}

// 求和并求平均值
func SumAndAverage(a []int) (int, float32) {
	sum := 0
	for _, value := range a {
		sum += value
	}

	return sum, float32(sum / len(a))
}
