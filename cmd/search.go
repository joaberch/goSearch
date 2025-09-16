package cmd

import (
	"fmt"
	"github.com/joaberch/goSearch/internal/model"
	"github.com/joaberch/goSearch/utils"
	"log"
	"os"
)

// Search looks up word in an inverted index built from the current working directory.
// The match behavior is controlled by mode (model.MatchMode). It prints progress messages,
// builds an index for the current directory, searches it, and displays the results.
// If retrieving the working directory fails the process is terminated via log.Fatal.
func Search(word string, mode model.MatchMode) {
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
