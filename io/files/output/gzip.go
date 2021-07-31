package main

import (
	"compress/gzip"
	"os"
)

func main() {
	file, err := os.Create("io/files/output/test.txt.gz")

	if err != nil {
		panic(err)
	}

	writer := gzip.NewWriter(file)

	writer.Header.Name = "text-gzip.txt"
	writer.Write([]byte("gzip.Writer sample\n"))
	writer.Close()
}
