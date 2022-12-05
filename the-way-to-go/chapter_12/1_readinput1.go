package main

import "fmt"

var (
	firstName, lastName, s, s1 string
	i                          int
	f                          float32
	input                      = "56.12 / 5212 / Go / s"
	format                     = "%f / %d / %s / %s"
	input1                     = "hh 2 4 6"
)

func main() {
	fmt.Println("Please enter your full name: ")
	//fmt.Scanln(&firstName, &lastName)
	// 从标准输入读取并格式化
	_, err := fmt.Scanf("%s %s", &firstName, &lastName)
	if err != nil {
		return
	}
	fmt.Printf("Hi %s %s \n", firstName, lastName)

	// 扫描参数字符串，将连续的以空格分隔的值存储到连续的参数中
	_, err = fmt.Sscanf(input, format, &f, &i, &s, &s1)
	if err != nil {
		return
	}
	fmt.Println("From the string we read: ", f, i, s, s1)

	// 扫描扫描参数字符串，将连续的以空格分隔的储存到连续的参数中。换行符也算作空格。它返回成功扫描的项目数。如果它少于比参数数少，Err将报告原因。
	_, err = fmt.Sscan(input1, &s, &s1)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("From the input1 we read: ", s, s1)
}

//Please enter your full name:
//shan la
//Hi shan la
//From the string we read:  56.12 5212 Go s
//From the input1 we read:  hh 2
