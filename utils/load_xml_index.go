package utils

import (
	"encoding/xml"
	"github.com/joaberch/goSearch/internal/model"
	"log"
	"os"
	"sort"
)

// LoadXMLIndex parses an XML index file and returns an InvertedIndex structure.
func LoadXMLIndex(file *os.File) model.InvertedIndex {
	var document model.IndexDocument
	decoder := xml.NewDecoder(file)
	err := decoder.Decode(&document)
	if err != nil {
		log.Fatal(err)
	}

	var index model.InvertedIndex
	for _, entry := range document.Entries {
		index.Entries = append(index.Entries, model.InvertedIndexEntry{
			Word:  entry.Word,
			Files: entry.Files,
		})
	}

	sort.Slice(index.Entries, func(i, j int) bool {
		return index.Entries[i].Word < index.Entries[j].Word
	})
	return index
}
