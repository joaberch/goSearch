package cmd

import (
	"fmt"
	"github.com/joaberch/goSearch/utils"
	"log"
	"os"
)

// Search performs a word lookup in the inverted index of the current directory
func Search(word string, mode string) {
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Indexing directory %s...\n", path)
	index := utils.Indexate(path)

	fmt.Printf("Searching for word %s...\n", word)
	results := utils.SearchInIndex(index, word, mode)
	utils.DisplayResult(results, word)
}
