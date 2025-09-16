package utils

import (
	"github.com/joaberch/goSearch/internal/model"
	"sort"
)

// CreateIndex builds an inverted index that maps each unique token (whitespace-separated)
// to the files and line numbers where it appears.
//
// The function tokenizes each line of each input file using strings.Fields (whitespace
// separation only), records the line indices from FileData.Content where each token
// occurs, removes duplicate line entries per file, and sorts line numbers in ascending
// order. The returned InvertedIndex contains sorted entries by word; each entry lists
// the files (by path) and the corresponding sorted, unique line indices.
//
// Parameters:
//   - files: slice of FileData to index. Line numbers are the indices of FileData.Content
//     (zero-based).
//
// Returns:
//   - model.InvertedIndex containing one entry per token. If the input is empty, an
//     empty InvertedIndex is returned.
//
// Notes:
//   - Tokenization is whitespace-based only; punctuation is not stripped or normalized.
//   - The function does not perform error handling or normalization (e.g., case folding).
func CreateIndex(files []model.FileData) model.InvertedIndex {
	tempIndex := make(map[string]map[string][]int) //Word string -> (file string -> lines []int)

	//Step 1 - Create a temporary index that store the word in the file and the lines where it can be found
	//If 2 entries have the same word it won't be the same entry because ["word"]["file1"]=[lines] and ["word"]["file2"]=[lines]
	for _, file := range files { //Foreach FileData
		for lineNumber, line := range file.Content { //Foreach content
			tokens := Normalize(line)

			for _, token := range tokens { //Foreach word
				if tempIndex[token] == nil { //If the word has never been processed, init the map
					tempIndex[token] = make(map[string][]int)
				}
				lines := tempIndex[token][file.Path]
				tempIndex[token][file.Path] = append(lines, lineNumber)
			}
		}
	}

	//Step 2 - Create the InvertedIndex and convert the temporary index in it
	var invertedIndex model.InvertedIndex
	for word, fileMap := range tempIndex {
		var fileMatches []model.FileMatch

		for path, lines := range fileMap {
			lines = RemoveDuplicates(lines)
			sort.Ints(lines) //To display the line in order, in case it has a lot of occurrences
			fileMatches = append(fileMatches, model.FileMatch{
				Name:  path,
				Lines: lines,
			})
		}

		invertedIndex.Entries = append(invertedIndex.Entries, model.InvertedIndexEntry{
			Word:  word,
			Files: fileMatches,
		})
	}

	sort.Slice(invertedIndex.Entries, func(i, j int) bool {
		return invertedIndex.Entries[i].Word < invertedIndex.Entries[j].Word
	})
	return invertedIndex
}
