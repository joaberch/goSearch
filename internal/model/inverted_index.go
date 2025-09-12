package model

// InvertedIndex maps the word with all the files that include it
type InvertedIndex map[string][]string
import "sort"


// ToXMLDocument converts an InvertedIndex to an IndexDocument
func (index *InvertedIndex) ToXMLDocument() IndexDocument {
	var entries []IndexEntry

	for _, entry := range index.Entries {
		xmlEntry := IndexEntry{
			Word:  entry.Word,
			Files: entry.Files,
		}
		entries = append(entries, xmlEntry)
	}
	sort.Slice(entries, func(i, j int) bool {
		return entries[i].Word < entries[j].Word
	})
	return IndexDocument{Entries: entries}
}
