package utils

import (
	"bufio"
	"compress/gzip"
	"io"
	"log"
	"os"
)

// CompressFile compress the file
func CompressFile(filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	//Create output file
	destPath := filePath + ".gz"
	dest, err := os.Create(destPath)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		err := dest.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	//Create writer
	writer := gzip.NewWriter(dest)
	defer func() {
		err := writer.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	//Copy source to dest
	_, err = io.Copy(writer, bufio.NewReader(file))
	if err != nil {
		log.Fatal(err)
	}
}
