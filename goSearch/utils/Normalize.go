package utils

import (
	"regexp"
	"strings"
)

// Normalization guideline :
// 1 - Everything in lowercase
// 2 - Remove punctuation
// 3 - Remove special character
// 4 - Remove words (the, is, and, etc) - not required but surely needed in that application
// 5 - Stemming
// 6 - Lemmatization

// Normalize string given and returns it
func Normalize(line string) string {
	//Step 1 - Put everything in lowercase
	line = strings.ToLower(line)

	//Step 2 & 3 - Remove punctuation and special character
	myRegexp, _ := regexp.Compile(`[^\w\s]`)
	line = myRegexp.ReplaceAllString(line, "")

	//Array with each word to remove useless one
	words := strings.Fields(line)

	//Step 4 - Remove words
	var filteredWords []string
	for _, word := range words {
		if CheckWordValidity(word) {
			filteredWords = append(filteredWords, word)
		}
	}

	//Step 5 - Stemming
	//TODO

	//Step 6 - Lemmatization
	//TODO

	return strings.Join(filteredWords, "; ")
}
