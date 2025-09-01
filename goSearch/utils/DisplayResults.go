package utils

import "fmt"

func DisplayResults(results map[string]bool, word string) {
	fmt.Printf("\nFound %d file(s) for \"%s\":\n", len(results), word)
	for path := range results {
		fmt.Printf("\t%s\n", path)
	}
}
