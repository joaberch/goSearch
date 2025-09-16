package utils

import (
	"log"
	"regexp"
	"strings"
)

// Normalization guideline:
// 1 - Everything in lowercase
// 2 - Remove punctuation
// 3 - Remove special character
// 4 - Remove words (the, is, and, etc.) - not required but surely needed in that application
// 5 - Stemming
// 6 - Lemmatization

// Normalize applies basic text preprocessing to a line: lowercasing, punctuation removal, and word filtering.
// Normalize lowercases, trims, strips punctuation/special characters, and returns
// the list of valid words extracted from the input line.
//
// Normalize performs basic text preprocessing: it converts the input to lower
// case, trims surrounding whitespace, removes characters that are not word
// characters or whitespace, and splits the result on whitespace. Each token is
// kept only if CheckWordValidity returns true. The function returns the slice
// of filtered words in the same order they appeared in the input. If the
// regular expression used for character removal fails to compile, the program
// will terminate via log.Fatal.
//
// Future enhancements may include stemming and lemmatization.
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
	//FUTURE (snowball?)

	//Step 6 - Lemmatization
	//FUTURE

	return filteredWords
}
