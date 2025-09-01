package utils

import (
	"goSearch/internal/model"
	"strings"
)

func SearchInIndex(index model.InvertedIndex, word string) []string {
	var results []string
	for key, paths := range index {
		if strings.Contains(strings.ToLower(key), strings.ToLower(word)) {
			results = append(results, paths...)
		}
	}
	return results
}
