package compresser

import (
	"bufio"
	"compress/gzip"
	"io"
	"log"
	"os"
)

func CompressXMLFile() {
	fileSource := "index.xml"
	fileDestination := "index.xml.gz"

	//Open source file
	source, err := os.Open(fileSource)
	if err != nil {
		log.Fatal(err)
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(source)

	//Create destination file
	destination, err := os.Create(fileDestination)
	if err != nil {
		log.Fatal(err)
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(destination)

	//Create writer
	w := gzip.NewWriter(destination)
	defer func(w *gzip.Writer) {
		err := w.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(w)

	//Copy content
	_, err = io.Copy(w, bufio.NewReader(source))
	if err != nil {
		log.Fatal(err)
	}
}
