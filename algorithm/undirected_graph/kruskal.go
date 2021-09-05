package main

import "sort"

type UnionFind struct {
	// データiが属する木の親番号。i==Parent[i]の時データiは木のノードである
	parent []int
	// 根ノードiに含まれるデータの数。iが根ノードでない場合は無意味な値となる
	sizes []int
}

func initUnionFind(n int) *UnionFind {
	u := new(UnionFind)
	u.parent = make([]int, n)
	for i := range u.parent {
		u.parent[i] = -1
	}
	return u
}

func (self *UnionFind) root(x int) int {
	if self.parent[x] == x {
		return x
	}

	// 経路圧縮
	self.parent[x] = self.root(self.parent[x])

	return self.parent[x]
}
func (self *UnionFind) unite(x, y int) bool {
	x = self.root(x)
	y = self.root(y)

	if x == y {
		return false
	}
	// yの木がxの木より大きくなるようにする
	if self.sizes[y] < self.sizes[x] {
		x, y = swap(x, y)
	}

	// unite
	self.parent[x] = y
	self.sizes[y] += self.sizes[x]
	return true
}

func (self *UnionFind) same(x, y int) bool {
	return self.root(x) == self.root(y)
}

func swap(x, y int) (int, int) {
	temp := x
	x = y
	y = temp
	return x, y
}

type Edge struct {
	a, b, cost int
}

type Graph struct {
	n     int    // 頂点の数
	edges []Edge //辺
}

func (self *Graph) kruskal() int {
	// コストが小さい順にソート
	sort.Slice(self.edges, func(i, j int) bool { return self.edges[i].cost < self.edges[j].cost })

	// UnionFind作成
	u := initUnionFind(self.n)
	min_cost := 0

	for i := 0; i < len(self.edges); i++ {
		e := self.edges[i]
		// 同じじゃない、つまり連結していない場合、連結を行う
		if !u.same(e.a, e.b) {
			// コストを足す(最小かどうかの判断はすでにソートしているため、必要ない)
			min_cost += e.cost
			u.unite(e.a, e.b)
		}
	}
	return min_cost
}
