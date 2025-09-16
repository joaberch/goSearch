package utils

import (
	"encoding/xml"
	"github.com/joaberch/goSearch/internal/model"
	"log"
	"os"
	"sort"
)

// LoadXMLIndex parses the provided XML index file into a model.InvertedIndex.
// It decodes an IndexDocument from file, converts each document entry into an
// InvertedIndexEntry, sorts entries by the Word field (ascending), and returns
// the resulting index. If XML decoding fails the function logs the error and
// terminates the program via log.Fatal.
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
