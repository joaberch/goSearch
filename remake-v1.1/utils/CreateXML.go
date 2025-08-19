package utils

import (
	"encoding/xml"
	"os"
	"search/internal/model"
)

// CreateXML creates an XML stocking the inverted index in a harcoded location - TODO: user choose/specific location
func CreateXML(InvertedIndex model.InvertedIndex) {
	//Step 1 - Create file
	file, err := os.Create("index.xml") //Currently create the index at the same location - TODO: name index from path
	if err != nil {
		panic(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	//Step 2 - Convert InvertedIndex to XML
	invIndex := ConvertInvertedIndexToXML(InvertedIndex)

	//Step 3 - Output XML to XML file
	encoder := xml.NewEncoder(file)
	encoder.Indent("", "\t")
	err = encoder.Encode(invIndex)
	if err != nil {
		return
	}
}
