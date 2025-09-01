package cmd

import (
	"fmt"
	"goSearch/utils"
	"os"
	"path/filepath"
)

func SearchWithIndex(args []string) {
	//Step 1 - Check args
	if len(args) < 2 {
		ShowHelp()
		return
	}
	word := args[0]
	indexName := args[1] + ".xml"

	//Step 2 - Get index path
	homedir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	indexPath := filepath.Join(homedir, "Desktop", "utils", "index", indexName)
	if _, err := os.Stat(indexPath); os.IsNotExist(err) {
		fmt.Printf("index %v does not exist\n", indexName)
		return
	}

	//Step 3 - Open path
	file, err := os.Open(indexPath)
	if err != nil {
		panic(err)
	}
	defer func(file *os.File) {
		err = file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	//Step 4 - Convert to InvertedIndex
	index := utils.LoadXMLIndex(file)

	//Step 5 - Search in InvertedIndex
	results := utils.SearchInIndex(index, word)

	//Step 6 - Output result
	utils.DisplayResults(results, word)
}
