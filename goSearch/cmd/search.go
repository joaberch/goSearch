package cmd

import (
	"fmt"
	"goSearch/utils"
	"os"
	"strings"
)

// Search return result of word in inverted index - TODO: can surely clean code more
func Search(word string) {
	//Step 1 - Indexate current path
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	fmt.Println("Indexing current directory...")
	index := utils.Indexate(path)

	//Step 2 - Search in index
	var results []string
	for key, paths := range index {
		if strings.Contains(strings.ToLower(key), strings.ToLower(word)) {
			results = append(results, paths...)
		}
	}

	// Step 3 - Display results
	fmt.Printf("Found %d file(s) for \"%s\":\n\n", len(results), word)
	for _, result := range results {
		fmt.Println(result)
	}
}
