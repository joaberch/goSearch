package model

// IndexEntry maps a word to the list of files where it appears.
type IndexEntry struct { //An index for a word with its file
	Word  string      `xml:"word,attr"`
	Files []FileMatch `xml:"files>file"`
}
