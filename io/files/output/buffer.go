package main

import (
	"bytes"
	"fmt"
	"io"
)

func main() {
	var buffer bytes.Buffer
	io.WriteString(&buffer, "os")
	fmt.Println(buffer.String())
}
