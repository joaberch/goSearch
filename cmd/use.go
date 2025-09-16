package cmd

import (
	"fmt"
	"github.com/joaberch/goSearch/internal/model"
	"github.com/joaberch/goSearch/utils"
	"log"
	"os"
	"path/filepath"
)

// SearchWithIndex searches for word in a compressed XML index file and displays matching results.
// 
// It expects compressedIndexName without extension; the function appends ".xml.gz" and looks for the file
// under the executable's directory at "index/<compressedIndexName>.xml.gz". If the compressed file is not
// found, it prints a not-found message and returns. The function decompresses the archive to a temporary
// XML file (which it removes before returning), loads the inverted index from that XML, performs the search
// using the provided MatchMode, and prints the results.
//
// Note: fatal errors (e.g., failing to determine the executable path, open the decompressed file, or clean up)
// are logged with log.Fatal which terminates the process.
//
// Parameters:
// - word: the search term.
// - compressedIndexName: base name of the compressed index (extension is added automatically).
// - mode: match mode that controls how terms are matched (exact, prefix, etc.).
func SearchWithIndex(word string, compressedIndexName string, mode model.MatchMode) {
	compressedIndexName += ".xml.gz"

	exePath, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	exeDir := filepath.Dir(exePath)
	compressedIndexPath := filepath.Join(exeDir, "index", compressedIndexName) //output path
	if _, err := os.Stat(compressedIndexPath); os.IsNotExist(err) {
		fmt.Printf("Index \"%s\" not found at \"%s\"\n", compressedIndexName, compressedIndexPath)
		return
	}

	indexPath := utils.Decompress(compressedIndexPath)
	defer func() {
		err = os.Remove(indexPath)
		if err != nil {
			log.Fatal(err)
		}
	}()

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
	results := utils.SearchInIndex(index, word, mode)
	utils.DisplayResult(results, word)
}
