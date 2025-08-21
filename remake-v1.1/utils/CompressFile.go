package utils

import (
	"bufio"
	"compress/gzip"
	"io"
	"os"
)

// CompressFile compress the file given with its path and its name
func CompressFile(path string, name string) {
	//Step 1 - Open source file
	source, err := os.Open(path + "\\" + name)
	if err != nil {
		panic(err)
	}
	defer func(source *os.File) {
		err := source.Close()
		if err != nil {
			panic(err)
		}
	}(source)

	//Step 2 - Create output file
	dest, err := os.Create(path + ".gz")
	if err != nil {
		panic(err)
	}
	defer func(f *os.File) {
		err := dest.Close()
		if err != nil {
			panic(err)
		}
	}(dest)

	//Step 3 - Create writer
	writer := gzip.NewWriter(dest)
	defer func(writer *gzip.Writer) {
		err := writer.Close()
		if err != nil {
			panic(err)
		}
	}(writer)

	//Step 4 - Copy source to dest
	_, err = io.Copy(writer, bufio.NewReader(source))
	if err != nil {
		panic(err)
	}
}
