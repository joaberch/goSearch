package utils

import (
	"github.com/joaberch/goSearch/internal/model"
	"log"
	"regexp"
	"strings"
)

// SearchInIndex returns a map of file paths where the given word appears in the inverted index.
// The search is case-insensitive and matches partial words.
func SearchInIndex(index model.InvertedIndex, word string, mode string) map[string][]int {
	var results = make(map[string][]int)
	lowerWord := strings.ToLower(word)

	var regex *regexp.Regexp
	var err error
	if mode == "regex" {
		regex, err = regexp.Compile(lowerWord)
		if err != nil {
			log.Fatal(err)
		}
	}

	for _, entry := range index.Entries {
		entry.Word = strings.ToLower(entry.Word)

		isMatching := false
		switch mode {
		case "contains":
			isMatching = strings.Contains(entry.Word, lowerWord)
		case "exact":
			isMatching = entry.Word == lowerWord
		case "regex":
			isMatching = regex.MatchString(entry.Word)
		}
		if isMatching {
			for _, file := range entry.Files {
				results[file.Name] = append(results[file.Name], file.Lines...)
			}
		}
	}
	return results
}
