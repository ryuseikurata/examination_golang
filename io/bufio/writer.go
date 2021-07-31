package main

import (
	"bufio"
	"os"
)

func main() {
	buffer := bufio.NewWriter(os.Stdout)

	buffer.WriteString("bufio.writer")
	buffer.Flush()
	buffer.WriteString("example \n")

	buffer.Flush()
}
