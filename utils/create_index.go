package utils

import (
	"github.com/joaberch/goSearch/internal/model"
)

// CreateIndex builds an inverted index mapping each unique word to the list of files where it appears.
func CreateIndex(files []model.FileData) model.InvertedIndex {
	index := make(model.InvertedIndex)
	for _, file := range files { //FUTURE: Should I create packages of x by x for goroutine?
		seen := make(map[string]bool)
		for _, token := range file.Content {
			if !seen[token] {
				index[token] = append(index[token], file.Path)
				seen[token] = true
			}
		}
	}
	return index
}
