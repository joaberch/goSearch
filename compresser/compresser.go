package compresser

import (
	"bufio"
	"compress/gzip"
	"io"
	"log"
	"os"
)

func DecompressGZtoXMLFile() {
	fileSource := "index.xml.gz"
	fileDestination := "index.xml"

	//open source file
	source, err := os.Open(fileSource)
	if err != nil {
		log.Fatal(err)
	}
	defer func(source *os.File) {
		err := source.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(source)

	//Create destination file
	destination, err := os.Create(fileDestination)
	if err != nil {
		log.Fatal(err)
	}
	defer func(destination *os.File) {
		err := destination.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(destination)

	//Create reader
	gzReader, err := gzip.NewReader(source)
	if err != nil {
		log.Fatal(err)
	}
	defer func(gzReader *gzip.Reader) {
		err := gzReader.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(gzReader)

	//Copy decompressed data to destination file
	_, err = io.Copy(destination, gzReader)
	if err != nil {
		log.Fatal(err)
	}
}

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

func DeleteXML() {
	err := os.Remove("index.xml")
	if err != nil {
		log.Fatal(err)
	}
}
