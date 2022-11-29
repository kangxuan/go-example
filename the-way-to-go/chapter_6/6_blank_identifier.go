package main

func main() {
	// 忽略返回值
	a, _, c := ThreeValues()
	print(a)
	print(c)
}

func ThreeValues() (int, float64, int) {
	return 1, 2.0, 3
}
