package main

import "strconv"

func main() {
	var lines []string
	max_value := -20000
	min_value := 20000

	values := intSliceFromStrSlice(lines)

	if err != nil {
		return
	}
}

func intSliceFromStrSlice(str_lines []string) ([]int, error) {
	var int_slice = make([]int, 0)
	for _, value := range str_lines {
		in_value, err := strconv.Atoi(value)
		if err != nil {
			return ([0], err)
		}
		int_slice = append(int_slice, in_value)
	}
	return int_slice
}
