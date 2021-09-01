package main

import "fmt"

func main() {
	array := []int{3, 4, 2, 7, 8, 4, 6, 7, 1}
	fmt.Println(array)
	array = bubble_sort(array)
	fmt.Println(array)
}

func bubble_sort(array []int) []int {
	for j := 0; j < len(array); j++ {
		for i := len(array) - 1; i > j; i-- {
			if array[i-1] > array[i] {
				array = swap(array, i-1, i)
			}
		}
	}

	return array
}

func swap(array []int, index1 int, index2 int) []int {
	temp := array[index1]
	array[index1] = array[index2]
	array[index2] = temp
	return array
}
