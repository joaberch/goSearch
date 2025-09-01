package utils

import "github.com/joaberch/goSearch/internal/model"

// CheckWordValidity returns true if the word should be included in the index.
func CheckWordValidity(word string) bool {
	if isInvalid, exists := model.InvalidWord[word]; exists {
		return !isInvalid
	}
	return true // Every other word is considered valid
}
