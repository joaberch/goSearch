package utils

import (
	"github.com/joaberch/goSearch/internal/model"
	"strings"
)

// SearchInIndex returns a map of file paths where the given word appears in the inverted index.
// The search is case-insensitive and matches partial words.
func SearchInIndex(index model.InvertedIndex, word string, mode string) map[string][]int {
	var results = make(map[string][]int)
	lowerWord := strings.ToLower(word)

	for _, entry := range index.Entries {
		entry.Word = strings.ToLower(entry.Word)

		if (mode == "contains" && strings.Contains(entry.Word, lowerWord)) ||
			(mode == "exact" && entry.Word == lowerWord) {
			for _, file := range entry.Files {
				results[file.Name] = append(results[file.Name], file.Lines...)
			}
		}
	}
	return results
}
