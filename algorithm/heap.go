package main

import (
	"fmt"
)

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
1 | 2 | 3 | 4 | 5 | 6 | 7 | 8
-----------------------------
1 |       |               |
**/
func min_heapify(heap Heap, i int) Heap {
	left_child := 2*i + 1
	right_child := 2*i + 2
	length := len(heap) - 1
	smallest := i

	// left_childがlengthより大きかったら => 葉なので停止
	// heap[smallest]がheap[left_child]より小さかったら、最小ノードを変更
	if left_child <= length && heap[smallest] > heap[left_child] {
		smallest = left_child
	}
	// right_childがlengthより大きかったら => 葉なので停止
	// heap[left_child]がheap[right_child]より大きかったら、最小ノードを変更
	if right_child <= length && heap[left_child] > heap[right_child] {
		smallest = right_child
	}
	// 最小ノードがiと同じでなかったら変更があるので、smallestとiの値を交換
	if smallest != i {
		// 交換処理
		temp := heap[i]
		heap[i] = heap[smallest]
		heap[smallest] = temp

		fmt.Printf("%d番目の%dと%d番目の%dが変わって、%dが出た。\n", i, temp, smallest, heap[smallest], heap)

		// smallestの場所から再帰処理
		min_heapify(heap, smallest)
	}
	return heap
}

func build_min_heap(heap Heap) {
	for i := len(heap) / 2; i >= 0; i-- {
		min_heapify(heap, i)
	}
}

func heap_sort(heapified_array Heap) []int {
	sorted_array := make([]int, 0, 0)
	heapified_array = min_heapify(heapified_array, 0)
	for i := 0; i <= len(heapified_array); i++ {
		lead := heapified_array[0]
		last := heapified_array[len(heapified_array)-1]
		heapified_array[0] = last
		heapified_array[len(heapified_array)-1] = lead

		sorted_array = append(sorted_array, last)
		fmt.Printf("%d\n", sorted_array)
		_, heapified_array, err := delete(heapified_array, len(heapified_array)-1)
		if err != nil {
			return []int{0}
		}
		build_min_heap(heapified_array)
	}
	return sorted_array
}

func main() {
	array := []int{4, 5, 6, 3, 2, 1, 9, 12, 8, -1, -2, 1, 4, 10, 5}
	fmt.Printf("%d\n", array)
	sorted_array := heap_sort(array)
	fmt.Printf("%d\n", sorted_array)
}

func delete(slice []int, i int) (int, []int, error) {
	ret := slice[i]
	if len(slice) < i || len(slice) < i {
		return 0, nil, fmt.Errorf("Error")
	}
	ans := make([]int, len(slice))
	copy(ans, slice)

	ans = append(slice[:i], slice[(i+1):]...)

	return ret, ans, nil
}
