package utils

import (
	"fmt"
	"sort"
)

// DisplayResult prints a summary of files that contain the given word.
// 
// It prints a header showing the number of files and the quoted search word, then
// lists each file path sorted alphabetically. For each file, the associated line
// numbers are deduplicated and sorted numerically before being displayed in the
// form: "\t<path> (lines: [<line1> <line2> ...])".
// 
// results maps a file path to the slice of line numbers where the word was found.
// word is the searched term shown in the header.
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
