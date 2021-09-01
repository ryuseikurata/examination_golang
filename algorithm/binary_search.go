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

// 再起的に降りて、目的の場所を見つけてノードを咲くする。削除した後にはdeterminの結果を埋め込む
func (n *Node) delete(v int) *Node {
	if n != nil {
		if n.value == v {
			if n.left == nil {
				return n.right
			} else if n.right == nil {
				return n.left
			} else {
				n.value = *&n.right.search_min().value
				n.right = n.right.delete_min()
			}
		} else if n.value < v {
			n.right = n.right.delete(v)
		} else {
			n.left = n.left.delete(v)
		}
	} else {
		return n
	}
	return n
}

// 最小値を取得
func (n *Node) search_min() *Node {
	for n.left != nil {
		n.left.search_min()
	}
	return n
}

func (n *Node) search_max() int {
	for n.right != nil {
		n.right.search_max()
	}
	return n.value
}

func (n *Node) delete_min() *Node {
	if n.left == nil {
		return n.right
	}
	n.left = n.left.delete_min()
	return n
}

func (n *Node) delete_max() *Node {
	if n.right == nil {
		return n.left
	}
	n.right = n.right.delete_max()
	return n
}

func (n *Node) String() string {
	if n == nil {
		return "nil"
	}
	return fmt.Sprintf("{value: %d, left: %v, right: %v}", n.value, n.left, n.right)
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

func main() {
	node := &Node{value: 15,
		left: &Node{
			value: 9, left: nil, right: &Node{
				value: 12, left: &Node{value: 10}, right: nil}},
		right: &Node{
			value: 19, left: &Node{
				value: 17, left: nil, right: nil},
			right: nil}}
	fmt.Println(node)
	node.delete(9)
	fmt.Println(node)
}
