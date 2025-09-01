package cmd

import (
	"fmt"
	"goSearch/utils"
	"os"
)

// Search return result of word in inverted index - TODO: can surely clean code more
func Search(args []string) {
	if len(args) < 1 {
		ShowHelp()
		return
	}
	word := args[0]
	//Step 1 - Indexate current path
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	fmt.Println("Indexing current directory...")
	index := utils.Indexate(path)

	//Step 2 - Search in index
	results := utils.SearchInIndex(index, word)

	// Step 3 - Display results
	utils.DisplayResults(results, word)
}
