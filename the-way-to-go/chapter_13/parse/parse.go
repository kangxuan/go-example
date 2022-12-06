package parse

import (
	"fmt"
	"strconv"
	"strings"
)

// ParseError 自定义一个错误结构
type ParseError struct {
	Index int    // 下标
	Word  string // 单词
	Err   error  // 错误
}

// ParseError的String方法
func (e *ParseError) String() string {
	return fmt.Sprintf("pkg parse: error parsing %q as int, index: %d", e.Word, e.Index)
}

// Parse 转换方法
func Parse(input string) (numbers []int, err error) {
	// defer和recover一起用，对panic进行捕获
	defer func() {
		if r := recover(); r != nil {
			var ok bool
			err, ok = r.(error)
			if !ok {
				err = fmt.Errorf("pkg: %v", r)
			}
		}
	}()

	// 将字符串分割成切片
	fields := strings.Fields(input)
	// 调用字段转换成数字
	numbers = fields2numbers(fields)
	return
}

// 字段转换成数字
func fields2numbers(fields []string) (numbers []int) {
	// 如果字段长度等于0，直接panic出异常
	if len(fields) == 0 {
		panic("no words to parse")
	}

	// 遍历字段数组并转换成数字
	for idx, field := range fields {
		num, err := strconv.Atoi(field)
		if err != nil {
			// 抛出ParseError会自动调用ParseError的String方法
			panic(&ParseError{idx, field, err})
		}
		// 将转换成功的数字追加的numbers切片中
		numbers = append(numbers, num)
	}

	return
}
