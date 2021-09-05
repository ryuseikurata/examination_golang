package main

func createTable(pattern string) []int {
	table := make([]int, len(pattern))
	table[0] = 0

	j := 0
	for i := 1; i < len(pattern); i++ {
		if pattern[i] == pattern[j] {
			table[i] = j
			j = j + 1
		} else {
			table[i] = j
			j = 0
		}
	}
	return table
}
func kmpSearch(target string, pattern string) int {
	table := createTable(pattern)
	i, p := 0, 0
	for {
		if i < len(target) && p < len(pattern) {

			// 文字が一致していれば次に進む
			if target[i] == pattern[p] {
				i++
				p++
			} else if p == 0 {
				// パターン先頭文字が不一致の場合、次の文字比較
				i++
			} else {
				// 先頭文字以外の不一致の場合、どの位置から再開するかを決定
				p = table[p]
			}

		} else {
			break
		}
	}
	if p == len(pattern) {
		return i - p
	}
	return -1
}
