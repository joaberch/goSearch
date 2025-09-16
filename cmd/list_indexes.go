package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

// ListIndexes prints saved index entries found in ~/Desktop/utils/index.
// 
// It looks up the current user's home directory, reads directory entries from
// Desktop/utils/index and prints either guidance messages when the index folder
// is missing or "No indexes found" when the folder is empty. For each entry it
// prints the file name and its modification time (format: "2006-01-02 15:04:05").
// The function exits the program via log.Fatal if it cannot determine the home
// directory or cannot retrieve metadata for an entry.
func ListIndexes() {
	exePath, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	exeDir := filepath.Dir(exePath)
	indexPath := filepath.Join(exeDir, "index") //output path
	
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
		info, err := entry.Info()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(" - ", entry.Name(), " - ", info.ModTime().Format("2006-01-02 15:04:05"))
	}
}
