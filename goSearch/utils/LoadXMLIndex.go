package utils

import (
	"encoding/xml"
	"goSearch/internal/model"
	"os"
)

func LoadXMLIndex(file *os.File) model.InvertedIndex {
	var document model.IndexDocument
	decoder := xml.NewDecoder(file)
	err := decoder.Decode(&document)
	if err != nil {
		panic(err)
	}

	index := make(model.InvertedIndex)
	for _, entry := range document.Entries {
		index[entry.Word] = entry.Files
	}
	return index
}
