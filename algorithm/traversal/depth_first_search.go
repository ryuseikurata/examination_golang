package main

import "fmt"

type Node struct {
	value int
	left  *Node
	right *Node
}

func DepthFirstSearch(node *Node) {
	if node != nil {
		DepthFirstSearch(node.left)
		DepthFirstSearch(node.right)
		fmt.Println(node)
	}
}
