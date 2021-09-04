package main

import (
	"container/list"
	"fmt"
)

type Graph struct {
	List []int
}

func BreadthFirstSearch(g *[]*Graph, n int, visited *[]int) {
	(*visited)[n] = 1

	q := list.New()
	q.PushBack(n)
	for {
		if q.Len() == 0 {
			break
		}

		n := q.Remove(q.Front())
		fmt.Printf("n:%d\n", n)
		fmt.Println(*visited)

		// 連結しているノードをforで行う
		for _, v := range (*g)[n.(int)].List {
			if (*visited)[v] == 0 {
				(*visited)[v] = 1
				fmt.Printf("\tn:%d->%d\n", n, v)
				q.PushBack(v)
			}
		}
	}
}

func main() {
	g := make([]*Graph, 0)
	g = append(g, &Graph{[]int{5, 6}}) // 0
	g = append(g, &Graph{[]int{3, 4}}) // 1
	g = append(g, &Graph{[]int{1}})    // 2
	g = append(g, &Graph{[]int{4}})    // 3
	g = append(g, &Graph{[]int{3}})    // 4
	g = append(g, &Graph{[]int{0}})    // 5
	g = append(g, &Graph{[]int{7}})    // 6
	g = append(g, &Graph{[]int{2}})    // 7

	visited := make([]int, len(g))
	BreadthFirstSearch(&g, 0, &visited)
}
