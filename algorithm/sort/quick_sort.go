package main

import "fmt"

func getPivot(array []int) int {
	if array[0] >= array[1] {
		return array[0]
	} else {
		return array[1]
	}
}

func quicksort(array []int) []int {
	if len(array) <= 1 {
		return array
	}

	pivot := getPivot(array)
	array0 := quicksort(getArrayGtPivot(array, pivot))
	array1 := quicksort(getArrayLtPivot(array, pivot))

	for i := 0; i < len(array0); i++ {
		array1 = append(array1, array0[i])
	}

	return array1
}

func getArrayGtPivot(array []int, pivot int) []int {
	gt_array := []int{}
	for i := 0; i < len(array); i++ {
		if array[i] >= pivot {
			gt_array = append(gt_array, array[i])
		}
	}
	return gt_array
}

func getArrayLtPivot(array []int, pivot int) []int {
	lt_array := []int{}
	for i := 0; i < len(array); i++ {
		if array[i] < pivot {
			lt_array = append(lt_array, array[i])
		}
	}
	return lt_array
}

func main() {
	array := []int{3, 4, 2, 7, 8, 4, 6, 7, 1}
	fmt.Println(array)
	array = quicksort(array)
	fmt.Println(array)
}
