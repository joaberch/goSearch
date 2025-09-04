package cmd

import (
	"fmt"
	"github.com/joaberch/goSearch/utils"
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
	compressedIndexName := args[1] + ".xml.gz"

	homedir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	compressedIndexPath := filepath.Join(homedir, "Desktop", "utils", "index", compressedIndexName)
	if _, err := os.Stat(compressedIndexPath); os.IsNotExist(err) {
		fmt.Printf("Index \"%s\" not found at \"%s\"\n", compressedIndexName, compressedIndexPath)
		return
	}

	indexPath := utils.Decompress(compressedIndexPath)
	defer func(name string) {
		err := os.Remove(name)
		if err != nil {
			log.Fatal(err)
		}
	}(indexPath)

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

	fmt.Printf("Searching for word %s in index file %s\n", word, compressedIndexPath)
	index := utils.LoadXMLIndex(file) //Convert IndexDocument to InvertedIndex
	results := utils.SearchInIndex(index, word)
	utils.DisplayResult(results, word)
}
