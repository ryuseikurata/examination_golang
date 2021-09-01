package main

import (
	"errors"
	"fmt"
)

// ノード
type Node struct {
	name  string  // ノード名
	edges []*Edge //次に移動できるエッジ
	done  bool    // 処理済みかを表すフラグ
	cost  int     // このノードに辿り着くのに必要だったコスト
	prev  *Node   // このノードに辿り着くのに使われたノード
}

type Edge struct {
	next *Node
	cost int
}

type DirectedGraph struct {
	nodes map[string]*Node
}

func NewNode(name string) *Node {
	// コストが-1は、そのノードまでの最短経路が未確定
	node := &Node{name, []*Edge{}, false, -1, nil}
	return node
}

func NewEdge(next *Node, cost int) *Edge {
	edge := *&Edge{next, cost}
	return &edge
}

func NewDirectedGraph() *DirectedGraph {
	return &DirectedGraph{
		map[string]*Node{},
	}
}

// ノードに次の接続先のエッジを追加
func (self *Node) AddEdge(edge *Edge) {
	self.edges = append(self.edges, edge)
}

func (self *DirectedGraph) Add(src, dst string, cost int) {
	// ノードがすでにある場合は追加しない
	// Map型は、第二引数で、要素が存在するかをboolで返す
	srcNode, ok := self.nodes[src]
	// ノードがなかったら追加
	if !ok {
		srcNode = NewNode(src)
		self.nodes[src] = srcNode
	}

	dstNode, ok := self.nodes[dst]

	if !ok {
		dstNode = NewNode(dst)
		self.nodes[dst] = dstNode
	}

	// ノードをEdgeで繋ぐ
	edge := NewEdge(dstNode, cost)
	srcNode.AddEdge(edge)
}

func (self *DirectedGraph) NextNode() (next *Node, err error) {
	// グラフに存在するノードを線形探索
	for _, node := range self.nodes {

		if node.done {
			continue
		}

		// costが-1の時、最短経路が判明してないので処理しない
		if node.cost == -1 {
			continue
		}

		// nextがnil => まだnextにnodeを入れてない => 最初の探索node
		// 問答無用でnodeを入れる
		if next == nil {
			next = node
		}

		// すでに見つかっているnode(next)よりもcostが小さいものがあれば、先にそちらを処理
		if next.cost > node.cost {
			next = node
		}
	}

	if next == nil {
		return nil, errors.New("Untreated node not found")
	}
	return
}
func (self *DirectedGraph) ShortestPath(start string, goal string) (ret []*Node, err error) {
	// start地点のノード取得
	startNode := self.nodes[start]

	// 初期設定(start地点のノードのcostを0にする)
	startNode.cost = 0

	for {
		// 次の処理対象のノードを取得する
		node, err := self.NextNode()

		if err != nil {
			return nil, errors.New("Goal not found")
		}

		//ゴールまで到達した
		if node.name == goal {
			break
		}

		// 取得したノードを処理する
		self.calc(node)
	}

	// ゴールから逆順にスタートまでたどっていく
	n := self.nodes[goal]
	for {
		ret = append(ret, n)
		if n.name == start {
			break
		}
		n = n.prev
	}

	return ret, nil
}

func (self *DirectedGraph) calc(node *Node) {
	// ノードにつながっているエッジを取得する
	for _, edge := range node.edges {
		nextNode := edge.next

		// すでに処理ずみのノードならスキップ
		if nextNode.done {
			continue
		}
		cost := nextNode.cost + edge.cost

		if nextNode.cost == -1 || cost < nextNode.cost {
			// 左は初期状態。右は、既に見つかっている経路よりもコストが小さいことを表している => sore
			// あれば処理中のノードを遷移元として記録
			nextNode.cost = cost
			nextNode.prev = node
		}
	}

	node.done = true
}

func main() {
	g := NewDirectedGraph()
	g.Add("s", "a", 2)
	g.Add("s", "b", 5)
	g.Add("a", "b", 2)
	g.Add("a", "c", 5)
	g.Add("b", "c", 4)
	g.Add("b", "d", 2)
	g.Add("c", "z", 7)
	g.Add("d", "c", 5)
	g.Add("d", "z", 2)

	// "s" ノードから "z" ノードへの最短経路を得る
	path, err := g.ShortestPath("c", "z")

	// 経路が見つからなければ終了
	if err != nil {
		fmt.Println("Goal not found")
		return
	}
	// 見つかった経路からノードとコストを表示する
	for _, node := range path {
		fmt.Printf("ノード: %v, コスト: %v\n", node.name, node.cost)
	}
}
