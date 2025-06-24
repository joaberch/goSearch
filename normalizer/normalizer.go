package normalizer

import (
	"regexp"
	"strings"
)

var stopWords = map[string]bool{
	"the": true, "is": true, "and": true, "a": true, "to": true, "in": true, "of": true, "their": true,
	"theirs": true, "de": true,
}

// Normalization guideline :
// 1 - Everything in lowercase
// 2 - Remove punctuation
// 3 - Remove special character
// 4 - Remove stop words (the, is, and, etc) - not required but surely needed in that application
// 5 - Stemming
// 6 - Lemmatization
func Normalize(line string) []string {
	//1 - Lowercase
	line = strings.ToLower(line)

	//2 & 3 - Punctuation & Special character
	regxp, _ := regexp.Compile(`[^\w\s]`)
	line = regxp.ReplaceAllString(line, "")

	//Split for each character
	words := strings.Fields(line)

	//4 - Remove stop words
	var filtered []string
	for _, word := range words {
		if !stopWords[word] {
			filtered = append(filtered, word)
		}
	}

	//5 - Stemming and Lemmatization

	return filtered
}
