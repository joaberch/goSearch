package model

type IndexEntry struct { //An index for a word with its file
	Word  string   `xml:"word,attr"`
	Files []string `xml:"files>file"`
}
