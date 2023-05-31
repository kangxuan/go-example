package main

import (
	"fmt"
	"github.com/willf/bloom"
)

const _1W = 10000

func main() {
	// 创建一个布隆过滤器，设置期望插入的元素数量和期望的误判率
	// 一般误判率越小，精确度越高，但对应的坑位增多以及哈希次数越多导致效率不高
	size := _1W * 100
	bf := bloom.NewWithEstimates(uint(size), 0.01) // 假设期望插入100W个元素，误判率为0.01

	// 向布隆过滤器中初始化加入100W个元素
	for i := 0; i < size; i++ {
		bf.Add([]byte(string(rune(i))))
	}

	// 故意取10万个不在过滤器里的值，看看有多少个会被认为在过滤器里，验证一下误判率
	count := 0
	for i := size + 1; i < size+(10*_1W); i++ {
		if bf.Test([]byte(string(rune(i)))) {
			fmt.Printf("被误判了:%d\n", i)
			count++
		}
	}
	fmt.Printf("有%d个元素被误判了", count)
}
