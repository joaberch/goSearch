package utils

import (
	"github.com/joaberch/goSearch/internal/model"
	"log"
	"regexp"
	"strings"
)

// SearchInIndex returns a map of file paths where the given word appears in the inverted index.
// Case-insensitive search. Mode "contains" does substring match; mode "exact" matches whole keys; mode "regex" matches regex expression.
func SearchInIndex(index model.InvertedIndex, word string, mode model.MatchMode) map[string][]int {
	var results = make(map[string][]int)
	lowerWord := strings.ToLower(word)

	var regex *regexp.Regexp
	var err error
	if mode == model.Regex {
		regex, err = regexp.Compile(lowerWord)
		if err != nil {
			log.Fatal(err)
		}
	}

	for _, entry := range index.Entries {
		entry.Word = strings.ToLower(entry.Word)

		isMatching := false
		switch mode {
		case model.Contains:
			isMatching = strings.Contains(entry.Word, lowerWord)
		case model.Exact:
			isMatching = entry.Word == lowerWord
		case model.Regex:
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
