package main

import "fmt"

type Node struct {
	value int
	left  *Node
	right *Node
}

func (n *Node) insert(v int) {
	for n == nil {
		if n.value <= v {
			n.right.insert(v)
		} else {
			n.left.insert(v)
		}
	}

	n = &Node{
		value: v,
		left:  nil,
		right: nil,
	}
	return
}

func main() {
	var node Node = Node{value: 15,
		left: &Node{
			value: 9, left: nil, right: &Node{
				value: 12, left: &Node{value: 10}, right: nil}},
		right: &Node{
			value: 19, left: &Node{
				value: 17, left: nil, right: nil},
			right: nil}}
	node.insert(10)
	fmt.Println(node)
}

func (n *Node) search(v int) *int {
	for n != nil {
		if n.value == v {
			return &n.value
		} else if n.value < v {
			return n.right.search(v)
		} else {
			return n.left.search(v)
		}
	}
	return nil
}
