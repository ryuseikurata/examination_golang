package main

type Heap []int
type IHeap interface {
	root() int
	parent(i int) int
	left(i int) int
	right(i int) int
}

func root(heap Heap) int {
	return heap[0]
}

func parent(heap Heap, i int) int {
	return heap[i/2-1]
}

func left(heap Heap, i int) int {
	return heap[2*i-1]
}

func right(heap Heap, i int) int {
	return heap[2*i]
}

/**
 そのノードと配下のノードがヒープの条件を満たしているか(ノードがその中で最小になっているか)
**/
func isHeap(heap Heap, i int) bool {
	n := heap[i-1]
	l := left(heap, i)
	r := right(heap, i)
	return n >= l && n >= r
}

/**
あるノードとその配下のノードがヒープの条件を満たすようにする操作
1.あるノードとその子ノードはヒープの条件を満たさない
2.子ノードとその配下のノードがヒープの条件を満たす時
**/
func min_heapify(heap Heap, i int) Heap {
	left := 2*i + 1
	right := 2*i + 2
	length := len(heap) - 1
	smallest := i

	if left <= length && heap[i] > heap[left] {
		smallest = left
	}
	if right <= length && heap[smallest] > heap[right] {
		smallest = right
	}
	if smallest != i {
		heap[i] = heap[smallest]
		heap[smallest] = heap[i]
		min_heapify(heap, smallest)
	}
	return heap
}
