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

// 再起的に降りて、目的の場所を見つけてノードを咲くおjする。削除した後にはdeterminの結果を埋め込む
func (n *Node) delete(v int) *Node {
	for n != nil {
		if n.value == v {
			// 削除ノードに子がない時
			if n.left == nil && n.right == nil {
				// 削除
				n = nil
				return n
			} else if n.left != nil && n.right == nil {
				// 削除ノードが一つの時(leftが存在) => leftを移動
				n = n.left
				return n
			} else if n.left == nil && n.right != nil {
				// 削除ノードが一つの時(rightが存在) => rightを移動
				n = n.right
				return n
			} else {
				// 削除ノードが二つの時
				//　左部分木の最大値 or 右部分木の最小値を移動
				// 今回は右部分木の最小値を移動させる
				// 1. 最小値があるnode取得
				// 2.
				tmp := n.right.search_min()
				n.value = tmp.value
				n.right = n.right.delete(tmp.value)
				return n
			}

		} else if n.value < v {
			return n.right.delete(v)
		} else {
			return n.left.delete(v)
		}
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
	node.delete(9)
	fmt.Println(node)
}
