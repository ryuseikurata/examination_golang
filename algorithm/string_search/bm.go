package main

import "math"

type Dict struct {
	char string
	i    int
}

func bmSearch(target, pattern string) int {
	table := createTable(pattern)
	i := len(pattern) - 1
	p := 0

	for {
		if i < len(target) {
			// パターン末尾に合わせる
			p = len(pattern) - 1

			for {
				if p >= 0 && i < len(target) {
					// 文字列が一致
					if target[i] == pattern[p] {
						i--
						p--
					} else {
						break
					}
				}
			}
			if p < 0 {
				return i + 1
			}

			shift1 := 0
			if val, ok := table[pattern[p]]; ok {
				shift1 = val
			} else {
				shift1 = len(pattern)
			}

			shift2 := len(pattern) - p
			i += int(math.Max(float64(shift1), float64(shift2)))
		}
		return -1
	}

}

func createTable(pattern string) map[byte]int {
	table := make(map[byte]int)
	for i := 0; i < len(pattern); i++ {
		// 同じ文字があっても、更新されるようになっている
		table[pattern[i]] = len(pattern) - i - 1
	}
	return table
}
