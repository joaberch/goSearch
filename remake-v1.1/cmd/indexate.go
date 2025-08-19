package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"search/internal/model"
	"search/utils"
)

func Indexate(path string) {
	//Step 1 - Get folder path
	var pathSelected string
	var err error
	var dir string

	if path == "-ui" { //if ui wanted
		pathSelected, err = utils.UISelectFolder()
	} else {
		dir, err = filepath.Abs(filepath.Dir(os.Args[0]))
		if err != nil {
			log.Fatal(err)
		}
		pathSelected = dir + "\\" + path
	}

	//Step 2 - create file tree and filter it (filter is inside CreateFileTree)
	tree := utils.CreateFileTree(pathSelected) //TODO: Check shortcut to prevent stack overflow - I think it's seen as a file, to check
	fmt.Println(tree)

	//Step 3 - Stream files (and normalize it in StreamFile)
	var contents []model.FileData
	for _, file := range tree.Children {
		streamRes := utils.StreamFile(file)
		contents = append(contents, streamRes)
	}

	//Step 4 - Save result in inverted index var
	invIndex := utils.CreateIndex(contents)

	//Step 5 - Save result in XML - TODO: user choose other file type? (json, etc)
	utils.CreateXML(invIndex, "index.xml")

	//Step 6 - Compress the XML file to store it

}
