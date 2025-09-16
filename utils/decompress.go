package utils

import (
	"compress/gzip"
	"io"
	"log"
	"os"
	"strings"
)

// Decompress decompresses the gzip file at gzPath and writes the decompressed data to a new file
// whose path is gzPath with the ".gz" suffix removed. It returns the path of the decompressed file.
// On any I/O or gzip errors the function calls log.Fatal, causing the program to exit.
func Decompress(gzPath string) string {
	zipFile, err := os.Open(gzPath)
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

	indexPath := strings.TrimSuffix(gzPath, ".gz")

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
