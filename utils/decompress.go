package utils

import (
	"compress/gzip"
	"io"
	"log"
	"os"
	"strings"
)

func Decompress(zipPath string) string {
	zipFile, err := os.Open(zipPath)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		err := zipFile.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	reader, err := gzip.NewReader(zipFile)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		err := reader.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	indexPath := strings.TrimSuffix(zipPath, ".gz")

	indexFile, err := os.Create(indexPath)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		err := indexFile.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	_, err = io.Copy(indexFile, reader)
	if err != nil {
		log.Fatal(err)
	}

	return indexPath
}
