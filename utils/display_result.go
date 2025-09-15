package utils

import (
	"fmt"
	"sort"
)

// DisplayResult prints the list of file paths where the word was found.
func DisplayResult(results map[string][]int, word string) {
	fmt.Printf("\nFound %d file(s) for %q:\n", len(results), word)

	sorted := make([]string, 0, len(results))
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
