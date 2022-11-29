package main

import "fmt"

func main() {
	var month = 10
	fmt.Printf("%d月是：%s", month, Season(month))
}

func Season(month int) string {
	switch month {
	case 3, 4, 5:
		return "春天"
	case 6, 7, 8:
		return "夏天"
	case 9, 10, 11:
		return "秋天"
	case 12, 1, 2:
		return "冬天"
	default:
		return "未知"
	}
}
