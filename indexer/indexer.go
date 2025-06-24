package indexer

import (
	"encoding/xml"
	"log"
	"os"
)

type XMLIndex struct {
	XMLName xml.Name   `xml:"index"`
	Entries []XMLEntry `xml:"entry"`
}

type XMLEntry struct {
	Word  string   `xml:"word,attr"`
	Files []string `xml:"file"`
}

type InvertedIndex map[string][]string

type DataFile struct {
	Path  string
	Token []string
}

func SaveIndexAsXml(index InvertedIndex, filename string) error {
	xmlData := ToXML(index)
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)

	encoder := xml.NewEncoder(file)
	encoder.Indent("", "  ")
	return encoder.Encode(xmlData)
}

func ToXML(index InvertedIndex) XMLIndex {
	var entries []XMLEntry
	for word, files := range index {
		entries = append(entries, XMLEntry{Word: word, Files: files})
	}
	return XMLIndex{Entries: entries}
}

func BuildInvertedIndex(files []DataFile) InvertedIndex {
	index := make(InvertedIndex) //Create an object
	for _, file := range files { //Foreach file
		seen := make(map[string]bool)      //Create a list of word added to the file
		for _, token := range file.Token { //Foreach word in the file
			if !seen[token] { //If the word doesn't already have the file path - could be removed to add a system of weight
				index[token] = append(index[token], file.Path) //Add the file path to the word
				seen[token] = true                             //Mark the word as processed
			}
		}
	}
	return index
}
