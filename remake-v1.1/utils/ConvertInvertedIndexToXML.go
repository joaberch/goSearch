package utils

import "search/internal/model"

func ConvertInvertedIndexToXML(index model.InvertedIndex) model.IndexDocument {
	var entries []model.IndexEntry
	for word, files := range index {
		entries = append(entries, model.IndexEntry{Word: word, Files: files})
	}
	return model.IndexDocument{Entries: entries}
}
