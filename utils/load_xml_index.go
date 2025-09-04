package utils

import (
	"encoding/xml"
	"github.com/joaberch/goSearch/internal/model"
	"log"
	"os"
)

// LoadXMLIndex parses an XML index file and returns an InvertedIndex structure.
func LoadXMLIndex(file *os.File) model.InvertedIndex {
	var document model.IndexDocument
	decoder := xml.NewDecoder(file)
	err := decoder.Decode(&document)
	if err != nil {
		log.Fatal(err)
	}

	index := make(model.InvertedIndex)
	for _, entry := range document.Entries {
		index[entry.Word] = entry.Files
	}
	return index
}
