package main

import "fmt"

type List []int

func (l List) Len() int {
	return len(l)
}

func (l *List) Append(val int) {
	*l = append(*l, val)
}

type Appender interface {
	Append(int)
}

func CountInto(a Appender, start int, end int) {
	for i := start; i <= end; i++ {
		a.Append(i)
	}
}

type Lener interface {
	Len() int
}

func LongEnough(l Lener) bool {
	return l.Len()*10 > 42
}

func main() {
	var lst List
	// LongEnough需要Lener，Lender定义在值上，可以调用
	if LongEnough(lst) {
		fmt.Printf("- lst is long enough\n")
	}

	//会报错，因为CountInto需要Appender，Appender定义在指针上
	//CountInto(lst, 1, 10)

	plst := new(List)
	// CountInto需要Appender，Appender定义在 指针上，plst就是一个指针类型
	CountInto(plst, 1, 10)
	// 在 plst 上调用 LongEnough 也是可以的，因为指针会被自动解引用。
	if LongEnough(plst) {
		fmt.Printf("- plst is long enough\n")
	}
}

//- plst is long enough
