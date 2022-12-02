package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("products2.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// 定义3个字符串切片
	var col1, col2, col3 []string
	for {
		var v1, v2, v3 string
		// 读取每行并通过空格分隔，并将每个字串复制给v1,v2,v3
		_, err := fmt.Fscanln(file, &v1, &v2, &v3)
		if err != nil {
			break
		}
		// 让最终列切片写入数据
		col1 = append(col1, v1)
		col2 = append(col2, v2)
		col3 = append(col3, v3)
	}

	fmt.Println(col1)
	fmt.Println(col2)
	fmt.Println(col3)
}

//[kang xuan]
//[1 2]
//[100 200]
