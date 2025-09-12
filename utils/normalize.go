package utils

import (
	"log"
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

// Normalize applies basic text preprocessing to a line: lowercasing, punctuation removal, and stopword filtering.
// Future enhancement may include stemming and lemmatization.
func Normalize(line string) []string {
	line = strings.ToLower(line)
	line = strings.TrimSpace(line)

	//Remove punctuation and special character
	myRegexp, err := regexp.Compile(`[^\w\s]`)
	if err != nil {
		log.Fatal(err)
	}
	line = myRegexp.ReplaceAllString(line, "")

	//Array with each word to remove useless one
	words := strings.Fields(line)

	var filteredWords []string
	for _, word := range words {
		if CheckWordValidity(word) {
			filteredWords = append(filteredWords, word)
		}
	}

	//Step 5 - Stemming
	//FUTURE

	//Step 6 - Lemmatization
	//FUTURE

	return filteredWords
}
