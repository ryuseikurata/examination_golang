package main

import "fmt"

type Node struct {
	value int
	left  *Node
	right *Node
}

func (n *Node) insert(v int) {
	if v < n.value {
		if n.left == nil {
			n.left = &Node{value: v, left: nil, right: nil}
		} else {
			n.left.insert(v)
		}
	} else {
		if n.right == nil {
			n.right = &Node{value: v, left: nil, right: nil}
		} else {
			n.right.insert(v)
		}
	}
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
	node.insert(8)
	fmt.Println(node.left.left.value)
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
