package main

import "fmt"

func quicksort(array []int) []int {
	if len(array) <= 1 {
		return array
	}

	left := []int{}
	right := []int{}

	pivot := array[0]

	for _, v := range array[1:] {
		if v >= pivot {
			right = append(right, v)
		} else {
			left = append(left, v)
		}
	}
	right = quicksort(right)
	left = quicksort(left)

	ret := append(left, pivot)
	ret = append(ret, right...)

	return ret
}

func main() {
	array := []int{3, 4, 2, 7, 8, 4, 6, 7, 1}
	fmt.Println(array)
	array = quicksort(array)
	fmt.Println(array)
}
