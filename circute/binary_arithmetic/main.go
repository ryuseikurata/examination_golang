package main

import "fmt"

type Binary []bool

// 5 / (-3)を実行したい
func main() {

	// 被除数 5 0101
	// 除数 -3 1101
	divisor := Binary{false, true, false, true}
	dividend := Binary{true, true, false, true}

	// 初期化
	n := len(dividend)
	q := []bool{false}
	d0 := shiftRight(divisor, n-1)
	d1 := shiftLeft(dividend, n-1)

	for i := 0; i < n-1; i++ {
		if isCodeEqual(d0, d1) || isCodeZero(d0) {
			d0 = d0 - d1
		} else {
			d0 = d0 + d1
		}

		if isCodeEqual(d0, d1) || isCodeZero(d0) {
			q = shiftRight(q, 1)
		} else {
			// Qの一番右のビットを0とする
			q[len(q)-1] = false
		}

		d1 = shiftRight(d1, 1)
		q = shiftLeft(q, 1)
	}

	if q[len(q)-1] == false {
		d0 = d0 + d1
	}
	fmt.Printf("商は%t", q)
	fmt.Printf("余りは%t", d0)

}

func shiftRight(binary Binary, shift_count int) []bool {
	slice := make([]bool, len(binary))
	for i := 0; i < shift_count; i++ {
		slice = append([]bool{false}, slice...)
	}
	return slice
}

func shiftLeft(binary Binary, shift_count int) []bool {
	slice := make([]bool, len(binary))
	for i := 0; i < shift_count; i++ {
		slice = append(slice, false)
	}
	return slice
}

func isCodeEqual(a Binary, b Binary) bool {
	return a[0] == b[0]
}

func isCodeZero(slice Binary) bool {
	for _, value := range slice {
		if value == true {
			return false
		}
	}
	return true
}
