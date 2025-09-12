package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func ListIndexes() {
	homedir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	indexPath := filepath.Join(homedir, "Desktop", "utils", "index")
	entries, err := os.ReadDir(indexPath)
	if err != nil {
		fmt.Println("No index folder found")
		fmt.Println("Use the `--save` option to save the index")
		return
	}

	if len(entries) == 0 {
		fmt.Println("No indexes found")
		return
	}
	fmt.Println("Available indexes:")
	for _, entry := range entries {
		fmt.Println(" - ", entry.Name())
	}
}
