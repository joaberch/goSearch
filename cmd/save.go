package cmd

import (
	"encoding/xml"
	"github.com/joaberch/goSearch/utils"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// SaveIndex saves the inverted index for the given path to a compressed XML file.
// 
// The function generates an inverted index for abs(path), writes it as an indented
// XML file under "<exeDir>/index/<basename>.xml", compresses that file to
// "<basename>.xml.gz", removes the uncompressed XML, and logs the final .gz path.
// It creates the output directory if missing. Any error during these steps causes
// immediate termination via log.Fatal.
func SaveIndex(path string) {
	exePath, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	exeDir := filepath.Dir(exePath)
	indexPath := filepath.Join(exeDir, "index") //output path
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
