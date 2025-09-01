package cmd

import (
	"encoding/xml"
	"github.com/joaberch/Go-LocalSearchEngine/utils"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// SaveIndex saves the inverted index into an XML file
func SaveIndex(path string) {
	homedir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	indexPath := filepath.Join(homedir, "Desktop", "utils", "index") //FUTURE: user choose output path?
	if _, err = os.Stat(indexPath); os.IsNotExist(err) {
		err = os.MkdirAll(indexPath, os.ModePerm)
		if err != nil {
			log.Fatal(err)
		}
	}

	filename := filepath.Base(path)
	if !strings.HasSuffix(filename, ".xml") {
		filename += ".xml"
	}
	indexFile := filepath.Join(indexPath, filename)
	file, err := os.Create(indexFile)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		err = file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	index := utils.Indexate(path)
	xmlIndex := utils.ConvertInvertedIndexToXML(index)

	encoder := xml.NewEncoder(file)
	encoder.Indent("", "\t")
	err = encoder.Encode(xmlIndex)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Saved index to %s\n", indexFile)
}
