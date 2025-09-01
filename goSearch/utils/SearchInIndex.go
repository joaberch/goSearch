package utils

import (
	"goSearch/internal/model"
	"strings"
)

func SearchInIndex(index model.InvertedIndex, word string) map[string]bool {
	var results = make(map[string]bool)
	for key, paths := range index {
		if strings.Contains(strings.ToLower(key), strings.ToLower(word)) {
			for _, path := range paths {
				results[path] = true
			}
		}
	}
	return results
}
