package cmd

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strings"
)

func Search(word string) {
	var files []string
	hasFound := false

	//Step 1 - Open source file
	source, err := os.Open("index.xml")
	if err != nil {
		panic(err)
	}
	defer func(source *os.File) {
		err := source.Close()
		if err != nil {
			panic(err)
		}
	}(source)

	//Step 2 - Parse file to find word
	decoder := xml.NewDecoder(source)

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
				hasFound = false              //reset foreach entry
				for _, attr := range t.Attr { //foreach attribute
					if strings.Contains(attr.Value, word) { //if the attribute is the word we are looking for - TODO: if contains
						hasFound = true
					}
				}
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
