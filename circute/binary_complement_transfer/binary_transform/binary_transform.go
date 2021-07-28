package binary_transform

// 1. 8 ~ 0までの10進数が来る
// 2. 3桁で返す
func Transform(decimal_number int) [4]int {
	if decimal_number == 8 {
		return [4]int{1, 0, 0, 0}
	} else if decimal_number == 7 {
		return [4]int{0, 1, 1, 1}
	} else if decimal_number == 6 {
		return [4]int{0, 1, 1, 0}
	} else if decimal_number == 5 {
		return [4]int{0, 1, 0, 1}
	} else if decimal_number == 4 {
		return [4]int{0, 1, 0, 0}
	} else if decimal_number == 3 {
		return [4]int{0, 0, 1, 1}
	} else if decimal_number == 2 {
		return [4]int{0, 0, 1, 0}
	} else if decimal_number == 1 {
		return [4]int{0, 0, 0, 1}
	} else if decimal_number == 0 {
		return [4]int{0, 0, 0, 1}
	} else {
		return [4]int{9, 9, 9, 9}
	}
}
