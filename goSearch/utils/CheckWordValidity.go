package utils

import "goSearch/internal/model"

func CheckWordValidity(word string) bool { //AI generated
	if isInvalid, exists := model.InvalidWord[word]; exists {
		return !isInvalid
	}
	return true // Every other word is considered valid
}
