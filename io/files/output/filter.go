package main

import (
	"io"
	"os"
)

func main() {
	file, err := os.Create("io/files/output/multiwriter.txt")

	if err != nil {
		panic(err)
	}

	writer := io.MultiWriter(file, os.Stdout)

	io.WriteString(writer, "io.MultiReader example\n")
}
