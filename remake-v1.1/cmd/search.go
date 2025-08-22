package cmd

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"search/utils"
)

// Search return result of word in inverted index - TODO: can surely clean code more
func Search(word string) {
	var files []string
	hasFound := false

	//Step 1 - Open source file
	decoder, source, _ := utils.OpenXMLFile("index.xml")
	defer func(source *os.File) {
		err := source.Close()
		if err != nil {
			panic(err)
		}
	}(source)

	//Step 2 - Parse file to find word
	for { //infinite loop
		t, err := decoder.Token() //parse line by line
		if err == io.EOF {        //If eof we exit the infinite loop
			break
		}
		if err != nil {
			panic(err)
		}

		switch t := t.(type) {
		case xml.StartElement:
			if t.Name.Local == "entry" { //if is an entry
				hasFound = utils.MatchEntry(t, word)
			} else if hasFound && t.Name.Local == "file" { //if we find the word
				var path string
				err := decoder.DecodeElement(&path, &t) //read all content from files
				if err != nil {
					panic(err)
				}
				files = append(files, path)
			}
		}
	}

	fmt.Printf("Found %d file(s) :\n\n", len(files))
	for _, file := range files {
		fmt.Println(file)
	}
}
