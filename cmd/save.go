package cmd

import (
	"encoding/xml"
	"github.com/joaberch/goSearch/utils"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// SaveIndex saves the inverted index into an XML file
func SaveIndex(path string) {
	baseDir, err := os.UserConfigDir()
	if err != nil {
		log.Fatal(err)
	}
	indexPath := filepath.Join(baseDir, "Desktop", "utils", "index") //output path
	if _, err = os.Stat(indexPath); os.IsNotExist(err) {
		err = os.MkdirAll(indexPath, os.ModePerm)
		if err != nil {
			log.Fatal(err)
		}
	}

	absPath, err := filepath.Abs(path)
	if err != nil {
		log.Fatal(err)
	}
	filename := filepath.Base(absPath)
	if !strings.HasSuffix(filename, ".xml") {
		filename += ".xml"
	}
	indexFile := filepath.Join(indexPath, filename)
	file, err := os.Create(indexFile)
	if err != nil {
		log.Fatal(err)
	}

	index := utils.Indexate(absPath)
	xmlIndex := index.ToXMLDocument()

	encoder := xml.NewEncoder(file)
	encoder.Indent("", "\t")
	err = encoder.Encode(xmlIndex)
	if err != nil {
		log.Fatal(err)
	}

	utils.CompressFile(indexFile)

	err = file.Close()
	if err != nil {
		log.Fatal(err)
	}

	err = os.Remove(indexFile)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Saved index to %s.gz\n", indexFile)
}
