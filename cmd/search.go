package cmd

import (
	"fmt"
	"goSearch/utils"
	"log"
	"os"
)

// Search performs a word lookup in the inverted index of the current directory
func Search(args []string) {
	if len(args) < 1 {
		ShowHelp()
		return
	}
	word := args[0]
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Indexing directory %s...\n", path)
	index := utils.Indexate(path)

	fmt.Printf("Searching for word %s...\n", word)
	results := utils.SearchInIndex(index, word)
	utils.DisplayResult(results, word)
}
