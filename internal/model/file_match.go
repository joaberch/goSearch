package model

// FileMatch represents a file found after a search with its name and the line where the word is found
type FileMatch struct {
	Name  string
	Lines []int
}
