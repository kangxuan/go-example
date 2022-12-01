package main

import (
	"fmt"
	"time"
)

type myTime struct {
	time.Time
}

func (t myTime) first3Chars() string {
	return t.Time.String()[0:3]
}

func main() {
	myT1 := myTime{time.Now()}
	// myT1.string()是调用time.Time类型的方法
	fmt.Printf("Full time now:%s\n", myT1.String())
	// 调佣自己定义的firstChars
	fmt.Println("The first3Chars:", myT1.first3Chars())
}

//Full time now:2022-12-01 10:35:50.211209 +0800 CST m=+0.000064975
//The first3Chars: 202
