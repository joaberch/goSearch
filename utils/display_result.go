package utils

import (
	"fmt"
	"sort"
)

// DisplayResult prints the list of file paths where the word was found.
func DisplayResult(results map[string]bool, word string) { //FUTURE: Filter/Sort result
	fmt.Printf("\nFound %d file(s) for \"%s\":\n", len(results), word)
	var sorted []string
	for path := range results {
		sorted = append(sorted, path)
	}
	sort.Strings(sorted)

	for _, path := range sorted {
		fmt.Printf("\t%s\n", path)
	}
}
