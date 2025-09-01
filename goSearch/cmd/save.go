package cmd

import (
	"encoding/xml"
	"goSearch/utils"
	"os"
	"path/filepath"
)

func SaveIndex(path string) {
	//Step 1 - Create folder if not exist
	//folder 'index' in Desktop/utils/index
	homedir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	indexPath := filepath.Join(homedir, "Desktop", "utils", "index")
	if _, err = os.Stat(indexPath); os.IsNotExist(err) {
		err = os.MkdirAll(indexPath, os.ModePerm)
		if err != nil {
			panic(err)
		}
	}

	//Step 2 - Create file
	filename := filepath.Base(path) + ".xml"
	indexFile := filepath.Join(indexPath, filename)
	file, err := os.Create(indexFile)
	if err != nil {
		panic(err)
	}
	defer func(file *os.File) {
		err = file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	//Step 3 - Generate index object
	index := utils.Indexate(path)
	//Step 4 - Generate XML Object
	xmlIndex := utils.ConvertInvertedIndexToXML(index)
	//Step 5 - Output the XML object into an XML file
	encoder := xml.NewEncoder(file)
	encoder.Indent("", "\t")
	err = encoder.Encode(xmlIndex)
	if err != nil {
		panic(err)
	}
}
