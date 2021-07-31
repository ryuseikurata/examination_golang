package main

/**
 * @description FX取引では、異なる国の通貨を交換することで為替さの利益を得ることができます。例えば、1ドル100円の時に1000ドル買い、価格変動により1どる108円になった時にうると、(108円 - 100円) * 1000ドル = 8000円の利益を得ることができます。
 * ある通貨について、時刻tにおける価格Rt(t=0,1,2...n-1)が入力として与えられるので、価格の差Rj-Ri(ただし、j > i)の最大値を求めてください。
 * @param　最初の行に整数nが与えられれます。続くn行に整数Rt(t=0,1,2...n-1)が順番に与えられます。
 * @returns 最大値を1行に出力
 * @limit 2<=n<=200,000
 * 1<=R_t<=10^9
go run rasen_book/part2/ALDS1_1_D/maximum-profit.go
5
1 2 3 4 5
*/

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	// スペースで文字列を区切る

	sc.Split(bufio.ScanLines)
	n := nextInt()
	values := getNumList()
	max_value := values[0]
	min_value := values[0]

	for i := 0; i < n; i++ {
		value := values[i]
		max_value = getMax(max_value, value-min_value)

		min_value = getMin(min_value, value)
	}
	fmt.Printf("%d", max_value)
}

func getMax(value_1 int, value_2 int) int {
	if value_1 >= value_2 {
		return value_1
	} else {
		return value_2
	}
}

func getMin(value_1 int, value_2 int) int {
	if value_1 >= value_2 {
		return value_2
	} else {
		return value_1
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

func getNumList() []int {
	sc.Scan()
	ret := sc.Text()
	arr := strings.Split(ret, " ")
	num_list := make([]int, 0, 0)
	for _, v := range arr {
		num, err := strconv.Atoi(v)

		if err != nil {
			panic(err)
		}
		num_list = append(num_list, num)
	}
	return num_list
}
