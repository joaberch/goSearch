package utils

import (
	"github.com/joaberch/goSearch/internal/model"
	"sort"
)

// ConvertInvertedIndexToXML transforms an InvertedIndex into an IndexDocument for XML serialization.
func ConvertInvertedIndexToXML(index model.InvertedIndex) model.IndexDocument { //Future: Sort by word for manual research and/or diff git
	var entries []model.IndexEntry
	for word, files := range index {
		entries = append(entries, model.IndexEntry{Word: word, Files: files})
	}

	sort.Slice(entries, func(i, j int) bool {
		return entries[i].Word < entries[j].Word
	})

	return model.IndexDocument{Entries: entries}
}
