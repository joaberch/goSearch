package main

import (
	"Go-LocalSearchEngine/indexer"
	"Go-LocalSearchEngine/streamer"
	"Go-LocalSearchEngine/treebuilder"
	"Go-LocalSearchEngine/uier"
	"bufio"
	"compress/gzip"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	//Step 1 - Get Files
	//Step 1.1 - UI Folder selection
	selected := uier.SelectFolder()

	//Step 1.2 Filter file, we don't want to process exe file or dll or iso
	res := treebuilder.GetFileTree(selected)

	//Step 2 - Read file in streaming and normalize content
	//Step 2.1 - Stream and normalize file content
	var filesData []indexer.DataFile
	for _, file := range res.ValidFiles {
		toName := streamer.Stream(file)                                            //Shortcut are seen as file and not as folder, prevent stackOverflow error
		filesData = append(filesData, indexer.DataFile{Path: file, Token: toName}) //Step 2.2 - Store normalized content (temporarily)
	}

	//Step 3 - Create Inverted index
	invertedIndex := indexer.BuildInvertedIndex(filesData)
	fmt.Println(invertedIndex)
	//Step 3.1 - Create XML (Be wary of depth and display the number of file to index)
	err := indexer.SaveIndexAsXml(invertedIndex, "index.xml")
	if err != nil {
		log.Fatal("Erreur lors de la sauvegarde du fichier XML", err)
	}

	//Step 4 - Save Index
	//Step 4.1 - Compress XML to store
	compressXMLFile()
}

func compressXMLFile() {
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
