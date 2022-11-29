package main

import "fmt"

func main() {
	str := "Google"
	// 将字符串转成字符切片
	sl := []byte(str)
	fmt.Printf("sl: %s\n", sl)

	// 通过借助其他数组反转
	var rev [100]byte
	j := 0
	for i := len(sl) - 1; i >= 0; i-- {
		rev[j] = sl[i]
		j++
	}
	revStr := string(rev[:])
	fmt.Println(revStr)

	// 直接交换
	str2 := "Google"
	sl2 := []byte(str2)
	for i, j := 0, len(sl2)-1; i < j; i, j = i+1, j-1 {
		sl2[i], sl2[j] = sl2[j], sl2[i]
	}
	fmt.Println(string(sl2))

	// 函数
	str3 := "你好，世界"
	str3 = reverse(str3)
	fmt.Println(str3)
}

// 通过rune可以处理Unicode编码的字符
func reverse(str string) string {
	sl := []rune(str)
	n, h := len(sl), len(sl)/2

	for i := 0; i < h; i++ {
		sl[i], sl[n-i-1] = sl[n-i-1], sl[i]
	}
	return string(sl)
}
