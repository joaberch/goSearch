package utils

import "fmt"

func DisplayResults(results []string, word string) {
	fmt.Printf("\nFound %d file(s) for \"%s\":\n", len(results), word)
	for _, result := range results {
		fmt.Println(result)
	}
}
