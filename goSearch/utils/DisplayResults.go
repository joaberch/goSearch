package utils

import "fmt"

// DisplayResults prints the list of file paths where the word was found.
func DisplayResults(results map[string]bool, word string) { //FUTURE: Filter/Sort result
	fmt.Printf("\nFound %d file(s) for \"%s\":\n", len(results), word)
	for path := range results {
		fmt.Printf("\t%s\n", path)
	}
}
