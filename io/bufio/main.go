package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// ファイルを用意
	f, _ := os.Open("io/bufio/text.txt")

	defer f.Close()

	// スキャナを用意
	sc := bufio.NewScanner(f)

	// EOFに当たるまでスキャンを繰り返す
	for sc.Scan() {
		line := sc.Text()
		fmt.Println(line)
	}
}
