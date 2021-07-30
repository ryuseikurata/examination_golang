package main

/**
 * @description FX取引では、異なる国の通貨を交換することで為替さの利益を得ることができます。例えば、1ドル100円の時に1000ドル買い、価格変動により1どる108円になった時にうると、(108円 - 100円) * 1000ドル = 8000円の利益を得ることができます。
 * ある通貨について、時刻tにおける価格Rt(t=0,1,2...n-1)が入力として与えられるので、価格の差Rj-Ri(ただし、j > i)の最大値を求めてください。
 * @param　最初の行に整数nが与えられれます。続くn行に整数Rt(t=0,1,2...n-1)が順番に与えられます。
 * @returns 最大値を1行に出力
 * @limit 2<=n<=200,000
 * 1<=R_t<=10^9
 */

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	// スペースで文字列を区切る
	sc.Split(bufio.ScanWords)
	n := nextInt()
	var values []int
	for i := 0; i < n; i++ {
		m := nextInt()
		values = append(values, m)
	}

	max_value := values[0]
	min_value := values[0]

	for _, value := range values {
		max_value = getMin(max_value, value-min_value)
		min_value = getMin(min_value, value)
	}

	fmt.Printf("%d", max_value)

	// for文を回す
	// max_valueを保持(max演算で現在のmax_valueとvalue - min_valueを比較)
	// min_valueを保持(min演算で現在のmin_valueとvalueを比較)
}

func intSliceFromStrSlice(str_lines []string) ([]int, error) {
	var int_slice = make([]int, 0)
	for _, value := range str_lines {
		in_value, err := strconv.Atoi(value)
		if err != nil {
			return []int{0}, err
		}
		int_slice = append(int_slice, in_value)
	}
	return int_slice, nil
}

func getMin(value_1 int, value_2 int) int {
	if value_1 >= value_2 {
		return value_1
	} else {
		return value_2
	}
}

// 文字列を読み込み戻り値をintで返す
func nextInt() int {
	sc.Scan()
	ret, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return ret
}
