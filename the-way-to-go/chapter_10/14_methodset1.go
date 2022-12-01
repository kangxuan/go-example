package main

import "fmt"

type List []int

func (lst List) len() int {
	return len(lst)
}

func (lst *List) Append(val int) {
	*lst = append(*lst, val)
}

func main() {
	lst := List{} // 值
	lst.Append(1)
	fmt.Printf("%v len is %d\n", lst, lst.len())

	lst2 := new(List) // 指针
	lst2.Append(2)
	fmt.Printf("%v len is %d\n", lst2, lst2.len())
}

//[1] len is 1
//&[2] len is 1
