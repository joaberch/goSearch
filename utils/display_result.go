package utils

import (
	"fmt"
	"sort"
)

// DisplayResult prints the list of file paths where the word was found.
func DisplayResult(results map[string][]int, word string) {
	fmt.Printf("\nFound %d file(s) for \"%s\":\n", len(results), word)

	var sorted []string
	for path := range results {
		sorted = append(sorted, path)
	}
	sort.Strings(sorted)

	for _, path := range sorted {
		lines := results[path]
		lines = RemoveDuplicates(lines)
		sort.Ints(lines)
		fmt.Printf("\t%s (lines: %v)\n", path, lines)
	}
}
