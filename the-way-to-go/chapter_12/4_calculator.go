package main

import (
	"bufio"
	"fmt"
	"os"
	"stack"
	"strconv"
)

func main() {
	// 创建一个读取器，并将其与标准输入绑定
	buf := bufio.NewReader(os.Stdin)
	// 申请一个栈
	calc1 := new(stack.Stack)

	fmt.Println("Give a number, an operator (+, -, *, /), or q to stop:")
	// 循环监听标准输入
	for {
		token, err := buf.ReadString('\n')
		if err != nil {
			fmt.Println("Input error !")
			os.Exit(1)
		}
		// 移除\n
		token = token[0 : len(token)-1]

		fmt.Printf("-- %s --\n", token)

		switch {
		case token == "q": // 当遇到q则停止
			fmt.Println("Calculator stopped")
			return
		case token >= "0" && token <= "999999": // 当遇到数字则将其推入到栈中
			i, _ := strconv.Atoi(token) // 字符串数字转换成整型
			calc1.Push(i)
		default:
			Calc(calc1, token)
		}
	}
}

// Calc 计算方法
func Calc(calc *stack.Stack, symbol string) {
	q, _ := calc.Pop()
	p, _ := calc.Pop()
	result := 0
	switch symbol {
	case "+":
		result = p.(int) + q.(int)
	case "-":
		result = p.(int) - q.(int)
	case "*":
		result = p.(int) * q.(int)
	case "/":
		result = p.(int) * q.(int)
	default:
		fmt.Println("无效的符号")
	}
	fmt.Printf("The result of %d %s %d = %d\n", p, symbol, q, result)
}

//Give a number, an operator (+, -, *, /), or q to stop:
//100
//-- 100 --
//200
//-- 200 --
//+
//-- + --
//The result of 100 + 200 = 300
//q
//-- q --
//Calculator stopped
