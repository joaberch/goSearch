package model

// InvertedIndexEntry maps the word with all the files that include it
type InvertedIndexEntry struct {
	Word  string
	Files []FileMatch
}
