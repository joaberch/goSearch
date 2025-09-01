package cmd

import (
	"fmt"
	"github.com/joaberch/Go-LocalSearchEngine/utils"
	"log"
	"os"
	"path/filepath"
)

// SearchWithIndex searches for a word in a saved XML index file.
func SearchWithIndex(args []string) {
	if len(args) < 2 {
		ShowHelp()
		return
	}
	word := args[0]
	indexName := args[1] + ".xml"

	homedir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	indexPath := filepath.Join(homedir, "Desktop", "utils", "index", indexName)
	if _, err := os.Stat(indexPath); os.IsNotExist(err) {
		fmt.Printf("Index \"%s\" not found at \"%s\"\n", indexName, indexPath)
		return
	}

	file, err := os.Open(indexPath)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		err = file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	fmt.Printf("Searching for word %s in index file %s\n", word, indexPath)
	index := utils.LoadXMLIndex(file) //Convert IndexDocument to InvertedIndex
	results := utils.SearchInIndex(index, word)
	utils.DisplayResult(results, word)
}
