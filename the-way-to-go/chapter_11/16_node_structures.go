package main

import "fmt"

type Node struct {
	left  *Node
	data  interface{}
	right *Node
}

func NewNode(left, right *Node) *Node {
	return &Node{left, nil, right}
}

func (p *Node) SetNode(data interface{}) {
	p.data = data
}

func main() {
	// 生成一个新节点
	root := NewNode(nil, nil)
	root.SetNode("root node")

	l := NewNode(nil, nil)
	l.SetNode("left node")

	r := NewNode(nil, nil)
	r.SetNode("right node")

	root.left = l
	root.right = r
	fmt.Printf("root : %v\n", root)
}

//root : &{0xc00008a040 root node 0xc00008a060}
