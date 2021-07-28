package main

import (
	"fmt"

	"./binary_transform"
)

func main() {
	var decimal_number_1 = 5
	var decimal_number_2 = 3

	var binary_1 = binary_transform.Transform(decimal_number_1)
	var binary_2 = binary_transform.Transform(decimal_number_2)

	fmt.Printf("%d", binary_1)
	fmt.Printf("%d", binary_2)
}

func addOne(i_lines int[]) {
	last_item := i_lines[len(i_lines) - 1]

	if last_item == 0 {
		swapEndItem(i_lines, 1)
	} else {
		swapEndItem(i_lines, 0)

	}
}

func swapEndItem(i_lines int[], num int) {

		// 末尾削除
		i_lines = i_lines[:len(i_lines) - 1]
		// 末尾追加
		i_lines = append(i_lines, num)
}
