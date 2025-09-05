package utils

import (
	"github.com/joaberch/goSearch/internal/model"
	"strings"
)

// SearchInIndex returns a map of file paths where the given word appears in the inverted index.
// The search is case-insensitive and matches partial words.
func SearchInIndex(index model.InvertedIndex, word string, mode string) map[string]bool {
	var results = make(map[string]bool)
	lowerWord := strings.ToLower(word)
	for key, paths := range index {
		keyLower := strings.ToLower(key)
		if (strings.Contains(keyLower, lowerWord) && mode == "contains") || //Future: regex
			(mode == "exact" && keyLower == lowerWord) {
			for _, path := range paths {
				results[path] = true
			}
		}
	}
	return results
}
