package utils

import (
	"github.com/joaberch/Go-LocalSearchEngine/internal/model"
	"strings"
)

// SearchInIndex returns a map of file paths where the given word appears in the inverted index.
// The search is case-insensitive and matches partial words.
func SearchInIndex(index model.InvertedIndex, word string) map[string]bool {
	var results = make(map[string]bool)
	lowerWord := strings.ToLower(word)
	for key, paths := range index {
		if strings.Contains(strings.ToLower(key), lowerWord) { //FUTURE : user choose if contains or if equals
			for _, path := range paths {
				results[path] = true
			}
		}
	}
	return results
}
