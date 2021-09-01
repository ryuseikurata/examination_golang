package main

import "fmt"

// dpは隣接行列
func Search(dp [][]int) ([][]int, error) {
	// 中間ノード
	for k := 0; k < len(dp); k++ {
		// 接続元ノード
		for i := 0; i < len(dp); i++ {
			// 接続先ノード
			for j := 0; j < len(dp); j++ {
				cost := dp[i][j]
				other := dp[i][k] + dp[k][j]
				if other < cost {
					cost = other
				}
				dp[i][j] = cost
			}
		}
	}

	for i := range dp {
		if dp[i][i] < 0 {
			return dp, fmt.Errorf("negative cycle")
		}
	}
	return dp, nil
}

func main() {
	// 隣接行列の参考
	// https://yttm-work.jp/algorithm/algorithm_0014.html
	dp := [][]int{{0, 2, 5, 1}, {99, 0, 2, 99}, {99, 99, 0, 99}, {99, 99, 2, 0}}
	Search(dp)
	fmt.Println(dp)
}
